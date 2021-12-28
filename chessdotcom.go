package chessdotcomgo

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ATTron/chessdotcom-go/util"
)

func processRespJSON(resp *http.Response) (string, error) {
	if resp.StatusCode == 200 {
		b, _ := ioutil.ReadAll(resp.Body)
		var jsonResp map[string]interface{}
		err := json.Unmarshal(b, &jsonResp)
		util.Check(err)
		returnResp, err := json.MarshalIndent(&jsonResp, "", "  ")
		util.Check(err)
		return string(returnResp), nil
	} else if resp.StatusCode == 404 {
		log.Println("Unable to find the request you are looking for")
		return "", util.ErrNotFound
	} else {
		log.Fatal("Could not return valid response from server . . .")
		return "", util.ErrBadType
	}
}

func processRespArray(resp *http.Response) (string, error) {
	if resp.StatusCode == 200 {
		b, _ := ioutil.ReadAll(resp.Body)
		var jsonResp []interface{}
		err := json.Unmarshal(b, &jsonResp)
		util.Check(err)
		returnResp, err := json.MarshalIndent(&jsonResp, "", "  ")
		util.Check(err)
		return string(returnResp), nil
	} else if resp.StatusCode == 404 {
		log.Println("Unable to find the request you are looking for")
		return "", util.ErrNotFound
	} else {
		log.Fatal("Could not return valid response from server . . .")
		return "", util.ErrBadType
	}
}
