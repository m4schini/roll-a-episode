package tvdb

import (
	"context"
	"github.com/antihax/optional"
	"log"
	"net/http"
	"os"
	"roll-episode/tvdb/auth"
	"roll-episode/tvdb/swagger"
	"strconv"
)

var tvdb *swagger.APIClient

func init() {
	apikey := os.Getenv("TVDB_APIKEY")
	if apikey == "" {
		log.Fatal("TVDB_APIKEY is missing")
	}

	pin := os.Getenv("TVDB_PIN")
	if pin == "" {
		log.Fatal("TVDB_PIN is missing")
	}

	var err error
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

	return results.Data, err
}

func GetShow(ctx context.Context, id string) (*swagger.SeriesExtendedRecord, error) {
	seriesBase, _, err := tvdb.SeriesApi.GetSeriesExtended(ctx, toFloat(id), nil)
	return seriesBase.Data, err
}

func toFloat(idStr string) float64 {
	n, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0
	}
	return float64(n)
}
