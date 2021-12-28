package chessdotcomgo

import (
	"net/http"
	"strconv"

	"github.com/ATTron/chessdotcom-go/util"
)

// GetTournament - Returns tournament details
func GetTournament(tournamentID string) []byte {
	reqURL := util.Join(util.TOURNAMENT_URL, tournamentID)
	resp, err := http.Get(reqURL)
	util.Check(err)

	defer resp.Body.Close()
	returnResp, err := processRespJSON(resp)
	util.Check(err)
	return returnResp
}

// GetTournamentRound - Returns tournament round details
func GetTournamentRound(tournamentID string, round int) []byte {
	reqURL := util.Join(util.TOURNAMENT_URL, tournamentID, strconv.Itoa(round))
	resp, err := http.Get(reqURL)
	util.Check(err)

	defer resp.Body.Close()
	returnResp, err := processRespJSON(resp)
	util.Check(err)
	return returnResp
}

// GetTournamentRoundGroup - Returns tournament round group details
func GetTournamentRoundGroup(tournamentID string, round int, group int) []byte {
	reqURL := util.Join(util.TOURNAMENT_URL, tournamentID, strconv.Itoa(round), "/", strconv.Itoa(group))
	resp, err := http.Get(reqURL)
	util.Check(err)

	defer resp.Body.Close()
	returnResp, err := processRespJSON(resp)
	util.Check(err)
	return returnResp
}
