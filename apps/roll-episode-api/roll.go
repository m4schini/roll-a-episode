package main

import "math/rand"

func Roll(seasons ...int) (season int, episode int) {
	l := len(seasons)
	season = rand.Intn(l)

	episodes := seasons[season]
	episode = rand.Intn(episodes)

	return season, episode
}
