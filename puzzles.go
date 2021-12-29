package chessdotcomgo

import (
	"net/http"

	"github.com/ATTron/chessdotcom-go/util"
)

// GetDailyPuzzle - Returns details about the daily puzzle
func GetDailyPuzzle() map[string]interface{} {
	reqURL := util.Join(util.PUZZLE_URL)
	resp, err := http.Get(reqURL)
	util.Check(err)

	defer resp.Body.Close()
	returnResp, err := processRespJSON(resp)
	util.Check(err)
	return returnResp
}

// GetRandomPuzzle - Returns details about a random puzzle
func GetRandomPuzzle() map[string]interface{} {
	reqURL := util.Join(util.PUZZLE_URL, "/random")
	resp, err := http.Get(reqURL)
	util.Check(err)

	defer resp.Body.Close()
	returnResp, err := processRespJSON(resp)
	util.Check(err)
	return returnResp
}
