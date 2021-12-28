package util

import (
	"errors"
	"log"
	"strings"
)

var BASE_URL = "https://api.chess.com/pub"
var PLAYER_URL = Join(BASE_URL, "/player/")
var CLUB_URL = Join(BASE_URL, "/club/")
var TOURNAMENT_URL = Join(BASE_URL, "/tournament/")
var MATCH_URL = Join(BASE_URL, "/matches/")
var COUNTRY_URL = Join(BASE_URL, "/country/")
var PUZZLE_URL = Join(BASE_URL, "/puzzle/")
var STREAMER_URL = Join(BASE_URL, "/streamers/")
var LEADERBOARD_URL = Join(BASE_URL, "/leaderboards/")

var ErrBadType = errors.New("unable to find api request of that type")
var ErrNotFound = errors.New("no results found for your request")

// Check - checks if err should be thrown
func Check(err error) {
	if err != nil {
		log.Fatal("An error has occured :: ", err)
	}
}

// Join - combine strings
func Join(strs ...string) string {
	var sb strings.Builder
	for _, q := range strs {
		sb.WriteString(q)
	}
	return sb.String()

}
