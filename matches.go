package chessdotcomgo

import (
	"net/http"
	"strconv"

	"github.com/ATTron/chessdotcom-go/util"
)

// GetDailyTeamMatch - Returns daily team match details
func GetDailyTeamMatch(matchID int) string {
	reqURL := util.Join(util.MATCH_URL, strconv.Itoa(matchID))
	resp, err := http.Get(reqURL)
	util.Check(err)

	defer resp.Body.Close()
	returnResp, err := processRespJSON(resp)
	util.Check(err)
	return returnResp
}

// GetDailyTeamMatchBoard - Returns daily team match board details
func GetDailyTeamMatchBoard(matchID int, board int) string {
	reqURL := util.Join(util.MATCH_URL, strconv.Itoa(matchID), "/", strconv.Itoa(board))
	resp, err := http.Get(reqURL)
	util.Check(err)

	defer resp.Body.Close()
	returnResp, err := processRespJSON(resp)
	util.Check(err)
	return returnResp
}

// GetLiveTeamMatch - Returns live team match details
func GetLiveTeamMatch(matchID int) string {
	reqURL := util.Join(util.MATCH_URL, "/live/", strconv.Itoa(matchID))
	resp, err := http.Get(reqURL)
	util.Check(err)

	defer resp.Body.Close()
	returnResp, err := processRespJSON(resp)
	util.Check(err)
	return returnResp
}

// GetLiveTeamMatchBoard - Returns live team match board details
func GetLiveTeamMatchBoard(matchID int, board int) string {
	reqURL := util.Join(util.MATCH_URL, "/live", strconv.Itoa(matchID), "/", strconv.Itoa(board))
	resp, err := http.Get(reqURL)
	util.Check(err)

	defer resp.Body.Close()
	returnResp, err := processRespJSON(resp)
	util.Check(err)
	return returnResp
}
