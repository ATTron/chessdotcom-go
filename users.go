package chessdotcomgo

import (
	"net/http"
	"strconv"

	"github.com/ATTron/chessdotcom-go/util"
)

type UserInfo struct {
	BasicInfo   string `json:"basic_info"`
	Stats       string `json:"stats"`
	OnlineStats string `json:"online"`
}

// GetTitledPlayers - Returns an array of all titled players of specific type (i.e. GM, WGM, IM, etc)
func GetTitledPlayers(title string) map[string]interface{} {
	reqURL := util.Join(util.BASE_URL, "/titled/", title)
	resp, err := http.Get(reqURL)
	util.Check(err)

	defer resp.Body.Close()
	returnResp, err := processRespJSON(resp)
	util.Check(err)
	return returnResp
}

// GetUser - Returns basic info about a specific user
func GetUser(username string) map[string]interface{} {
	reqURL := util.Join(util.PLAYER_URL, username)
	resp, err := http.Get(reqURL)
	util.Check(err)

	defer resp.Body.Close()
	returnResp, err := processRespJSON(resp)
	util.Check(err)
	return returnResp
}

// GetUserStats - Returns statistics about a specific user
func GetUserStats(username string) map[string]interface{} {
	reqURL := util.Join(util.PLAYER_URL, username, "/stats")
	resp, err := http.Get(reqURL)
	util.Check(err)

	defer resp.Body.Close()
	returnResp, err := processRespJSON(resp)
	util.Check(err)
	return returnResp
}

// GetPlayerStatus (Currently Broken) - Returns T/F of players online status
func GetPlayerStatus(username string) map[string]interface{} {
	reqURL := util.Join(util.PLAYER_URL, username, "/is-online")
	resp, err := http.Get(reqURL)
	util.Check(err)

	defer resp.Body.Close()
	returnResp, err := processRespJSON(resp)
	util.Check(err)
	return returnResp
}

// GetAllData (Currently not working)- Return all user info in one object
func GetAllData(username string) UserInfo {
	UI := UserInfo{}
	// UI.BasicInfo = GetUser(username)
	// UI.Stats = GetUserStats(username)
	// currently online status is not working so omitting it here
	// UI.OnlineStats = GetPlayerStatus(username)

	return UI
}

// GetPlayerGames - Return current user games
func GetPlayerGames(username string) map[string]interface{} {
	reqURL := util.Join(util.PLAYER_URL, username, "/games")
	resp, err := http.Get(reqURL)
	util.Check(err)

	defer resp.Body.Close()
	returnResp, err := processRespJSON(resp)
	util.Check(err)
	return returnResp
}

// GetPlayerGamesToMove - Return player who has to move
func GetPlayerGamesToMove(username string) map[string]interface{} {
	reqURL := util.Join(util.PLAYER_URL, username, "/games/to-move")
	resp, err := http.Get(reqURL)
	util.Check(err)

	defer resp.Body.Close()
	returnResp, err := processRespJSON(resp)
	util.Check(err)
	return returnResp
}

// GetPlayerGamesMonthly - Return games the player has played in the past month
func GetPlayerGamesMonthly(username string) map[string]interface{} {
	reqURL := util.Join(util.PLAYER_URL, username, "/games/archives")
	resp, err := http.Get(reqURL)
	util.Check(err)

	defer resp.Body.Close()
	returnResp, err := processRespJSON(resp)
	util.Check(err)
	return returnResp
}

// GetPlayerArchive - Return games the player has played in a specific month
func GetPlayerGamesArchive(username string, year int, month int) map[string]interface{} {
	reqURL := util.Join(util.PLAYER_URL, username, "/games/", strconv.Itoa(year), "/", strconv.Itoa(month))
	resp, err := http.Get(reqURL)
	util.Check(err)

	defer resp.Body.Close()
	returnResp, err := processRespJSON(resp)
	util.Check(err)
	return returnResp
}

// GetPlayerGamesArchivePGN (Currently broken) - Return games the player has played in a specific month
func GetPlayerGamesArchivePGN(username string, year int, month int) map[string]interface{} {
	reqURL := util.Join(util.PLAYER_URL, username, "/games/", strconv.Itoa(year), "/", strconv.Itoa(month), "/pgn")
	resp, err := http.Get(reqURL)
	util.Check(err)

	defer resp.Body.Close()
	returnResp, err := processRespJSON(resp)
	util.Check(err)
	return returnResp
}

// GetPlayerMatches - Returns list of team matches the player is currently registered in
func GetPlayerMatches(username string) map[string]interface{} {
	reqURL := util.Join(util.PLAYER_URL, username, "/matches")
	resp, err := http.Get(reqURL)
	util.Check(err)

	defer resp.Body.Close()
	returnResp, err := processRespJSON(resp)
	util.Check(err)
	return returnResp
}

// GetPlayerTournaments
func GetPlayerTournaments(username string) map[string]interface{} {
	reqURL := util.Join(util.PLAYER_URL, username, "/tournaments")
	resp, err := http.Get(reqURL)
	util.Check(err)

	defer resp.Body.Close()
	returnResp, err := processRespJSON(resp)
	util.Check(err)
	return returnResp
}
