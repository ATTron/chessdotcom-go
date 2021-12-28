package chessdotcomgo

import (
	"net/http"

	"github.com/ATTron/chessdotcom-go/util"
)

// GetStreamers - Returns details about chess.com streamers
func GetStreamers() []byte {
	reqURL := util.Join(util.STREAMER_URL)
	resp, err := http.Get(reqURL)
	util.Check(err)

	defer resp.Body.Close()
	returnResp, err := processRespJSON(resp)
	util.Check(err)
	return returnResp
}

// GetLeaderboard - Returns details about the chess.com leaderboard
func GetLeaderboard() []byte {
	reqURL := util.Join(util.LEADERBOARD_URL)
	resp, err := http.Get(reqURL)
	util.Check(err)

	defer resp.Body.Close()
	returnResp, err := processRespJSON(resp)
	util.Check(err)
	return returnResp
}
