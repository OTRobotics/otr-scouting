package otrscouting

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const (
	FRC_TOKEN = ""
	TBA_TOKEN = ""
)

func frc_request(path string) string {
	req, err := http.NewRequest("POST", "https://frc-api.firstinspires.org/v2.0/"+path, nil)
	if err != nil {
		return ""
	}
	req.Header.Add("Authorization", "Basic "+FRC_TOKEN)
	req.Header.Add("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	return string(respBody)
}

func tba_request(path string) string {
	req, err := http.NewRequest("POST", "https://www.thebluealliance.com/api/v3/"+path, nil)
	if err != nil {
		return ""
	}
	req.Header.Add("X-TBA-Auth-Key", TBA_TOKEN)
	req.Header.Add("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	return string(respBody)
}

func frcapiGetEvent(c *gin.Context, event string, year int) EventRoot {
	var root EventRoot
	var jsonData map[string]interface{}

	resp := tba_request("/event/" + strconv.Itoa(year) + event)
	err := json.Unmarshal([]byte(resp), &jsonData)

	if err != nil {
		return EventRoot{}
	}

	root.EventId = strings.ToLower(event)
	root.Year = year
	root.EventName = jsonData["name"].(string)
	root.EventDate = "Week " + jsonData["week"].(string)
	root.EventCode = strconv.Itoa(year) + "_" + root.EventId
	root.FRCEvents = strings.ToUpper(event)

	return root
}
