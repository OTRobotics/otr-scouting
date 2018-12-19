package otrscouting

import (
	json2 "encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UploadData struct {
	PageTitle    string
	Uploaded     bool
	Unauthorized bool
}

type MatchUpload struct {
	MatchId   string `json:"match_id"`
	Level     string `json:"comp_level"`
	EventId   string `json:"event_id"`
	Year      int    `json:"year"`
	MatchData struct {
		Red1  RobotTemplate `json:"red_1"`
		Red2  RobotTemplate `json:"red_2"`
		Red3  RobotTemplate `json:"red_3"`
		Blue1 RobotTemplate `json:"blue_1"`
		Blue2 RobotTemplate `json:"blue_2"`
		Blue3 RobotTemplate `json:"blue_3"`
	} `json:"match_data"`
}

func (m *MatchUpload) toMatchTemplate() MatchTemplate {
	var mt = MatchTemplate{}

	mt.MatchId = strconv.Itoa(m.Year) + "_" + m.EventId + "_" + string(m.Level[0]) + m.MatchId
	mt.MatchNumber = m.MatchId
	mt.Blue1 = m.MatchData.Blue1
	mt.Blue2 = m.MatchData.Blue2
	mt.Blue3 = m.MatchData.Blue3

	mt.Red1 = m.MatchData.Red1
	mt.Red2 = m.MatchData.Red2
	mt.Red3 = m.MatchData.Red3

	return mt
}

func (m MatchUpload) MatchCode() string {
	return strconv.Itoa(m.Year) + "_" + m.EventId + "_" + m.MatchId
}

func GinDataTypeHandler(c *gin.Context) {
	json := getMatch(c, "2018_onwat_q1")
	c.JSON(200, json)
}

func GinUploadHandler(c *gin.Context) {
	tmpl := GetPageTemplate("upload.html", c)
	data := UploadData{
		PageTitle:    "Upload Event Info",
		Uploaded:     false,
		Unauthorized: true,
	}
	if c.Request.Method == "GET" {
		tmpl.Execute(c.Writer, data)
	} else if c.Request.Method == "POST" {
		data.Uploaded = true
		// Extract JSON
		pass := c.PostForm("pass")
		if pass == "OTR3474" {
			data.Unauthorized = false
		}
		if data.Unauthorized {
			tmpl.Execute(c.Writer, data)
			return
		}
		rawJson := c.PostForm("json")
		var matches []MatchUpload
		tmpl.Execute(c.Writer, data)
		if err := json2.Unmarshal([]byte(rawJson), &matches); err != nil {
			fmt.Fprint(c.Writer, err)
		}

		uploadMatchToDatastore(c, matches)
	}
}

func GinManualUploadHandler(c *gin.Context) {

}
