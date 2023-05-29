package main

import "math/rand"

func Roll(seasons ...int) (season int, episode int) {
	l := len(seasons)
	if l <= 0 {
		log.Warn("seasons less than zero leads to panic")
		return 0, 0
	}
	season = rand.Intn(l)

	episodes := seasons[season]
	episode = rand.Intn(episodes)

	return season, episode
}
