package chessdotcomgo

import (
	"fmt"
	"net/http"

	"github.com/ATTron/chessdotcom-go/util"
)

type UserInfo struct {
	BasicInfo   string `json:"basic_info"`
	Stats       string `json:"stats"`
	OnlineStats string `json:"online"`
}

// GetUser - Returns basic info about a specific user
func GetUser(username string) string {
	reqURL := util.Join(util.BASE_URL, "/player/", username)
	resp, err := http.Get(reqURL)
	util.Check(err)

	defer resp.Body.Close()
	returnResp, err := processRespJSON(resp)
	util.Check(err)
	return returnResp
}

// GetUserStats - Returns statistics about a specific user
func GetUserStats(username string) string {
	// reqURL := util.BASE_URL + "/trains/" + strconv.Itoa(trainNum)
	reqURL := util.Join(util.BASE_URL, "/player/", username, "/stats")
	fmt.Println(reqURL)
	resp, err := http.Get(reqURL)
	util.Check(err)

	defer resp.Body.Close()
	returnResp, err := processRespJSON(resp)
	util.Check(err)
	return returnResp
}

// GetTitledPlayers - Returns an array of all titled players of specific type (i.e. GM, WGM, IM, etc)
func GetTitledPlayers(title string) string {
	reqURL := util.Join(util.BASE_URL, "/titled/", title)
	resp, err := http.Get(reqURL)
	util.Check(err)

	defer resp.Body.Close()
	returnResp, err := processRespJSON(resp)
	util.Check(err)
	return returnResp
}

// GetPlayerStatus (Currently Broken) - Returns T/F of players online status
func GetPlayerStatus(username string) string {
	reqURL := util.Join(util.BASE_URL, "/player/", username, "/is-online")
	fmt.Println(reqURL)
	resp, err := http.Get(reqURL)
	util.Check(err)

	defer resp.Body.Close()
	returnResp, err := processRespJSON(resp)
	util.Check(err)
	return returnResp
}

// GetAllData - Return all user info in one object
func GetAllData(username string) UserInfo {
	UI := UserInfo{}
	UI.BasicInfo = GetUser(username)
	UI.Stats = GetUserStats(username)
	// currently online status is not working so omitting it here
	// UI.OnlineStats = GetPlayerStatus(username)

	return UI
}
