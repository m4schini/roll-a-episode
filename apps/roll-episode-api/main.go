package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/m4schini/logger"
	"io"
	"math"
	"net/http"
	"os"
	"roll-episode/cache"
	"roll-episode/s3"
	"roll-episode/tvdb"
	"roll-episode/tvdb/swagger"
	"strconv"
	"sync"
	"time"
)

var log = logger.Named("main").Sugar()

var imgCache cache.Map
var imgCacheMu sync.Mutex

func main() {
	bucketName := os.Getenv("BUCKET_NAME")
	endpoint := os.Getenv("BUCKET_ENDPOINT")
	accessKey := os.Getenv("BUCKET_ACCESS_KEY")
	accessSecret := os.Getenv("BUCKET_ACCESS_SECRET")

	bucket, err := s3.NewMinioBucket(context.TODO(), bucketName, endpoint, accessKey, accessSecret, true)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.New()
	r.Use(GinLogger(), gin.Recovery())
	r.GET("/img", func(c *gin.Context) {
		log := log.With("path", "/img", "sender", c.Request.RemoteAddr)
		defer log.Info("completed request")
		href := c.Request.URL.Query().Get("href")
		hrefB64 := base64.StdEncoding.EncodeToString([]byte(href))

		resp, err := http.Get(fmt.Sprintf("https://%v/%v/%v", endpoint, bucketName, hrefB64))
		if err == nil && resp.StatusCode == http.StatusOK {
			log.Info("returning cached data for " + href)
			c.Header("Cache-Control", "max-age=604800")
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Content-Type", "image/jpeg")
			_, err := io.Copy(c.Writer, resp.Body)
			resp.Body.Close()
			if err != nil {
				log.Error(c.AbortWithError(http.StatusInternalServerError, err))
			}
			return
		}
		resp.Body.Close()

		response, err := http.Get(href)
		if err != nil {
			log.Error(c.AbortWithError(http.StatusInternalServerError, err))
			return
		}
		defer response.Body.Close()

		buf, err := io.ReadAll(response.Body)
		if err != nil {
			log.Error(c.AbortWithError(http.StatusInternalServerError, err))
			return
		}

		log.Info("caching")
		err = bucket.Put(c, hrefB64, "image/jpeg", int64(len(buf)), bytes.NewReader(buf))
		if err != nil {
			log.Warnf("caching of %v failed: %v", hrefB64, err)
		}
		c.Header("Cache-Control", "max-age=604800")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Data(http.StatusOK, "image/jpeg", buf)
	})
	r.GET("/show", func(c *gin.Context) {
		log := log.With("path", "/show", "sender", c.Request.RemoteAddr)
		defer log.Info("completed request")
		id := c.Query("id")

		record, err := tvdb.GetShow(c, id)
		if err != nil {
			log.Error(c.AbortWithError(http.StatusInternalServerError, err))
			return
		}

		c.Header("Cache-Control", "max-age=604800")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusOK, record)
	})
	r.GET("/search", func(c *gin.Context) {
		log := log.With("path", "/search", "sender", c.Request.RemoteAddr)
		defer log.Info("completed request")
		query := c.Query("q")
		log.Debug("query: ", query)

		limitStr := c.Query("limit")
		limit, err := strconv.ParseInt(limitStr, 10, 64)
		if err != nil {
			limit = 3
		}

		shows, err := tvdb.FindShow(c, query)
		if err != nil {
			log.Error(c.AbortWithError(http.StatusInternalServerError, err))
			return
		}

		shows = shows[:int(math.Min(float64(len(shows)), float64(limit)))]

		c.Header("Cache-Control", "max-age=604800")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusOK, shows)
	})
	r.GET("/roll", func(c *gin.Context) {
		log := log.With("path", "/roll", "sender", c.Request.RemoteAddr)
		defer log.Info("completed request")
		id := c.Query("id")

		seasons, records := tvdb.GetSeasons(c, id)
		season, episode := Roll(seasons...)
		var rolledEpisode *swagger.EpisodeBaseRecord
		for _, record := range records {
			if int(record.SeasonNumber) == season && int(record.Number) == episode {
				rolledEpisode = &record
				break
			}
		}

		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusOK, rolledEpisode)
	})
	log.Fatal(r.Run("0.0.0.0:" + os.Getenv("PORT")))
}

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.String()

		// Process request
		c.Next()

		log.Infof(`[%v] %v "%v" | %v | %v (%v)`,
			c.Writer.Status(),
			c.Request.Method,
			path,
			time.Since(start), //TODO prometheus timer
			c.Request.RemoteAddr,
			c.GetHeader("Content-Type"))
	}
}
