package main

import (
	"encoding/json"
	"github.com/m4schini/logger"
	"io"
	"math"
	"net/http"
	"os"
	"roll-episode/cache"
	"roll-episode/tvdb"
	"roll-episode/tvdb/swagger"
	"strconv"
	"sync"
)

var log = logger.Named("main").Sugar()

var resCache cache.Map
var resCacheMu sync.Mutex

var imgCache cache.Map
var imgCacheMu sync.Mutex

func main() {
	http.HandleFunc("/img", func(w http.ResponseWriter, r *http.Request) {
		href := r.URL.Query().Get("href")

		imgCacheMu.Lock()
		defer imgCacheMu.Unlock()
		cached, isCached := imgCache.Get(href)
		if isCached {
			log.Info("CACHED: " + href)
			w.Header().Set("Content-Type", "image/jpeg")
			w.Header().Set("Cache-Control", "max-age=604800")
			w.Write(cached)
			w.WriteHeader(http.StatusOK)
			return
		}

		response, err := http.Get(href)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		buf, err := io.ReadAll(response.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		log.Info("CACHING: " + href)
		imgCache.Set(href, buf)

		w.Header().Set("Content-Type", "image/jpeg")
		w.Header().Set("Cache-Control", "max-age=604800")
		w.Write(buf)
		w.WriteHeader(http.StatusOK)
	})
	http.HandleFunc("/show", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")

		record, err := tvdb.GetShow(r.Context(), id)
		if err != nil {
			log.Error("ERR", err)
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		marshal, err := json.Marshal(record)
		if err != nil {
			log.Error("ERR", err)
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Cache-Control", "max-age=604800")
		w.Header().Set("content-type", "application/json")
		w.Write(marshal)
		w.WriteHeader(http.StatusOK)
	})
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")

		resCacheMu.Lock()
		defer resCacheMu.Unlock()
		cached, isCached := resCache.Get(query)
		if isCached {
			log.Info("CACHED: " + query)
			w.Header().Set("content-type", "application/json")
			w.Write(cached)
			w.WriteHeader(http.StatusOK)
			return
		}

		limitStr := r.URL.Query().Get("limit")
		limit, err := strconv.ParseInt(limitStr, 10, 64)
		if err != nil {
			limit = 3
		}
		shows, err := tvdb.FindShow(r.Context(), query)
		if err != nil {
			log.Info("ERR", err)
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		shows = shows[:int(math.Min(float64(len(shows)), float64(limit)))]

		b, err := json.Marshal(shows)
		if err != nil {
			log.Info("ERR", err)
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		resCache.Set(query, b)

		log.Info("CACHING: " + query)
		w.Header().Set("Cache-Control", "max-age=604800")
		w.Header().Set("content-type", "application/json")
		w.Write(b)
		w.WriteHeader(http.StatusOK)
	})
	http.HandleFunc("/roll", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")

		seasons, records := tvdb.GetSeasons(r.Context(), id)

		season, episode := Roll(seasons...)
		var rolledEpisode *swagger.EpisodeBaseRecord
		for _, record := range records {
			if int(record.SeasonNumber) == season && int(record.Number) == episode {
				rolledEpisode = &record
				break
			}
		}

		out, err := json.Marshal(rolledEpisode)
		if err != nil {
			log.Error("ERR", err)
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("content-type", "application/json")
		w.Write(out)
		w.WriteHeader(http.StatusOK)
	})
	log.Fatalln(http.ListenAndServe("0.0.0.0:"+os.Getenv("PORT"), nil))
}
