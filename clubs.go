package chessdotcomgo

import (
	"net/http"

	"github.com/ATTron/chessdotcom-go/util"
)

// GetClub - Returns details about a club
func GetClub(clubID string) []byte {
	reqURL := util.Join(util.CLUB_URL, clubID)
	resp, err := http.Get(reqURL)
	util.Check(err)

	defer resp.Body.Close()
	returnResp, err := processRespJSON(resp)
	util.Check(err)
	return returnResp
}

// GetClubMembers - Returns details a specific club members
func GetClubMembers(clubID string) []byte {
	reqURL := util.Join(util.CLUB_URL, clubID, "/members")
	resp, err := http.Get(reqURL)
	util.Check(err)

	defer resp.Body.Close()
	returnResp, err := processRespJSON(resp)
	util.Check(err)
	return returnResp
}

// GetClubMatches - Returns club match details
func GetClubMatches(clubID string) []byte {
	reqURL := util.Join(util.CLUB_URL, clubID, "/matches")
	resp, err := http.Get(reqURL)
	util.Check(err)

	defer resp.Body.Close()
	returnResp, err := processRespJSON(resp)
	util.Check(err)
	return returnResp
}
