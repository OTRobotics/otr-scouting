package otrscouting

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type LandingPageTemplate struct {
	PageTitle string
	Items []LandingPageData
}

type LandingPageData struct {
	Name string
	Url string
}

func GinTeamsHandler(c *gin.Context){
	c.JSON(200, gin.H{"teams": "Not implemented yet."})

}

func GinMatchesHandler(c *gin.Context){
	var data LandingPageTemplate
	currentYear := time.Now().Year()
	tmpl := GetPageTemplate("landing.html", c)
	events := GetYearEvents(time.Now().Year(), c)
	data.PageTitle=strconv.Itoa(currentYear) + " Matches"
	for _, v := range events {
		var eventTemp LandingPageData

		matches := getEventMatches(c, v.EventCode)

		for _, m := range matches {
			eventTemp.Url = "/match/" + m.MatchId
			eventTemp.Name = v.EventName + " - " + m.FriendlyName()
			data.Items = append(data.Items, eventTemp)
		}
	}

	tmpl.Execute(c.Writer, data)
}

func GinEventsHandler(c *gin.Context){
	var data LandingPageTemplate
	currentYear := time.Now().Year()
	tmpl := GetPageTemplate("landing.html", c)
	events := GetYearEvents(time.Now().Year(), c)
	data.PageTitle=strconv.Itoa(currentYear) + " Events"
	for _, v := range events {
		var eventTemp LandingPageData

		eventTemp.Name = v.EventName
		eventTemp.Url = "/event/" +v.EventCode
		data.Items = append(data.Items, eventTemp)
	}

	tmpl.Execute(c.Writer, data)
}