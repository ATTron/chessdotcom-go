package chessdotcomgo

import (
	"net/http"
	"strconv"

	"github.com/ATTron/chessdotcom-go/util"
)

// GetCountry - Returns details about specific country
func GetCountry(countryISO int) []byte {
	reqURL := util.Join(util.COUNTRY_URL, strconv.Itoa(countryISO))
	resp, err := http.Get(reqURL)
	util.Check(err)

	defer resp.Body.Close()
	returnResp, err := processRespJSON(resp)
	util.Check(err)
	return returnResp
}

// GetCountryPlayers - Returns details about specific country's players
func GetCountryPlayers(countryISO int) []byte {
	reqURL := util.Join(util.COUNTRY_URL, strconv.Itoa(countryISO), "/players")
	resp, err := http.Get(reqURL)
	util.Check(err)

	defer resp.Body.Close()
	returnResp, err := processRespJSON(resp)
	util.Check(err)
	return returnResp
}

// GetCountryClubs - Returns details about specific country's clubs
func GetCountryClubs(countryISO int) []byte {
	reqURL := util.Join(util.COUNTRY_URL, strconv.Itoa(countryISO), "/clubs")
	resp, err := http.Get(reqURL)
	util.Check(err)

	defer resp.Body.Close()
	returnResp, err := processRespJSON(resp)
	util.Check(err)
	return returnResp
}
