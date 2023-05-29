package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/m4schini/logger"
	"io"
	"math"
	"net/http"
	"os"
	"roll-episode/s3"
	"roll-episode/tvdb"
	"roll-episode/tvdb/swagger"
	"strconv"
	"time"
)

var log = logger.Named("main").Sugar()

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
		c.Header("Access-Control-Allow-Origin", "*")

		var rolledEpisode *swagger.EpisodeBaseRecord
		var i = 0
		for rolledEpisode == nil {
			seasons, records := tvdb.GetSeasons(c, id)
			season, episode := Roll(seasons...)
			for _, record := range records {
				if int(record.SeasonNumber) == season && int(record.Number) == episode {
					rolledEpisode = &record
					break
				}
			}
			log.Debug("Roll attempt ", i)
			i += 1
		}

		if rolledEpisode == nil {
			c.JSON(http.StatusNotFound, struct{}{})
			return
		}
		rolledEpisode.Image, _ = tvdb.UseS3(c, rolledEpisode.Image)

		j, err := json.Marshal(rolledEpisode)
		if err == nil {
			log.Info("rolled ", string(j))
		}
		c.JSON(http.StatusOK, rolledEpisode)
	})
	log.Fatal(r.Run("0.0.0.0:" + os.Getenv("PORT")))
}

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// log request
		log.Infow(fmt.Sprintf(`Handeled request %v '%v'`, c.Request.Method, c.Request.URL.String()),
			"status", c.Writer.Status(),
			"duration", time.Since(start).String(),
			"client", c.Request.RemoteAddr,
			"contentType", c.GetHeader("Content-Type"))
	}
}
