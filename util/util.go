package util

import (
	"errors"
	"log"
	"strings"
)

var BASE_URL = "https://api.chess.com/pub"

var ErrBadType = errors.New("unable to find api request of that type")
var ErrNotFound = errors.New("no results found for your request")

// Check -- checks if err should be thrown
func Check(err error) {
	if err != nil {
		log.Fatal("An error has occured! :: ", err)
	}
}

// Join -- combine strings
func Join(strs ...string) string {
	var sb strings.Builder
	for _, q := range strs {
		sb.WriteString(q)
	}
	return sb.String()

}
