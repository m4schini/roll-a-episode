package tvdb

import (
	"context"
	"encoding/base64"
	"errors"
	"github.com/antihax/optional"
	"github.com/m4schini/logger"
	"net/http"
	"os"
	"roll-episode/s3"
	"roll-episode/tvdb/auth"
	"roll-episode/tvdb/swagger"
	"strconv"
	"sync"
)

var log = logger.Named("tvdb").Sugar()

var tvdb *swagger.APIClient

var s3Bucket *s3.MinioBucket

func init() {
	apikey := os.Getenv("TVDB_APIKEY")
	if apikey == "" {
		log.Fatal("TVDB_APIKEY is missing")
	}

	pin := os.Getenv("TVDB_PIN")
	if pin == "" {
		log.Fatal("TVDB_PIN is missing")
	}

	bucketName := os.Getenv("BUCKET_NAME")
	endpoint := os.Getenv("BUCKET_ENDPOINT")
	accessKey := os.Getenv("BUCKET_ACCESS_KEY")
	accessSecret := os.Getenv("BUCKET_ACCESS_SECRET")

	bucket, err := s3.NewMinioBucket(context.TODO(), bucketName, endpoint, accessKey, accessSecret, true)
	if err != nil {
		log.Error(err)
	} else {
		s3Bucket = bucket
	}

	cfg := swagger.NewConfiguration()
	cfg.HTTPClient = http.DefaultClient
	cfg.HTTPClient.Transport, err = auth.NewAuth(
		cfg.BasePath,
		apikey,
		pin,
		http.DefaultTransport)
	if err != nil {
		panic(err)
	}
	tvdb = swagger.NewAPIClient(cfg)
}

func GetSeasons(ctx context.Context, idStr string) ([]int, []swagger.EpisodeBaseRecord) {
	n, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return []int{}, []swagger.EpisodeBaseRecord{}
	}
	id := float64(n)

	series, _, err := tvdb.SeriesApi.GetSeriesEpisodes(ctx, 0, id, "official", nil)
	if err != nil {
		return []int{}, []swagger.EpisodeBaseRecord{}
	}
	var seriesEpisodeCount = make([]int, 0)
	for _, episode := range series.Data.Episodes {
		if episode.SeasonNumber == 0 {
			continue
		}

		if len(seriesEpisodeCount) < int(episode.SeasonNumber) {
			seriesEpisodeCount = append(seriesEpisodeCount, 0)
		}

		seriesEpisodeCount[episode.SeasonNumber-1] = int(episode.Number)
		//fmt.Printf("S=%02v E=%02v: %v\n", episode.SeasonNumber, episode.Number, episode.Name)
	}

	return seriesEpisodeCount, series.Data.Episodes
}

func FindShow(ctx context.Context, query string) ([]swagger.SearchResult, error) {
	results, _, err := tvdb.SearchApi.GetSearchResults(ctx, &swagger.SearchApiGetSearchResultsOpts{
		Q: optional.NewString(query),
	})

	var wg sync.WaitGroup
	l := len(results.Data)
	wg.Add(l)
	newImageUrls := make([]string, l)
	for i, result := range results.Data {
		i := i
		url := result.ImageUrl
		go func() {
			newUrl, _ := UseS3(ctx, url)
			newImageUrls[i] = newUrl
			wg.Done()
			log.Debugw("Tried to use s3 instead of tvdb", "old", url, "new", newUrl)
		}()
	}
	wg.Wait()

	for i, _ := range results.Data {
		results.Data[i].ImageUrl = newImageUrls[i]
	}

	return results.Data, err
}

func GetShow(ctx context.Context, id string) (*swagger.SeriesExtendedRecord, error) {
	seriesBase, _, err := tvdb.SeriesApi.GetSeriesExtended(ctx, toFloat(id), nil)

	seriesBase.Data.Image, err = UseS3(ctx, seriesBase.Data.Image)
	return seriesBase.Data, err
}

func toFloat(idStr string) float64 {
	n, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0
	}
	return float64(n)
}

func UseS3(ctx context.Context, orignalUrl string) (string, error) {
	if s3Bucket == nil {
		return orignalUrl, errors.New("not initialized")
	}
	urlB64 := base64.StdEncoding.EncodeToString([]byte(orignalUrl))
	s3Url := s3Bucket.URL(urlB64)
	resp, err := http.Head(s3Url)
	if err == nil && resp.StatusCode == http.StatusOK {
		return s3Url, nil
	}

	resp, err = http.Get(orignalUrl)
	if err != nil {
		return orignalUrl, err
	}

	contentLengthStr := resp.Header.Get("Content-Length")
	contentLength, err := strconv.ParseInt(contentLengthStr, 10, 64)
	if err != nil {
		return orignalUrl, err
	}

	err = s3Bucket.Put(ctx, urlB64, "image/jpeg", contentLength, resp.Body)
	resp.Body.Close()
	if err != nil {
		return orignalUrl, err
	}

	return s3Url, nil
}
