package otrscouting

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// Changes per year, will need to be updated
type RobotTemplate struct {
	Team    int `json:"team"`
	PowerUp PowerUpRobot
}

// FORMAT: /match/:matchnumber
func GinMatchHandler(c *gin.Context) {
	name := c.Param("matchnumber")
	// Check match number/format against event/match
	// 2018_event_match
	split := strings.Split(name, "_")
	if len(split) < 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Match Number Format"})
	}
	year := split[0]
	event := split[1]

	eventT := getEvent(c, year+"_"+event)

	mTemp := getMatch(c, name)

	tmpl := GetPageTemplate("match.html", c)
	tmpl.Execute(c.Writer, mTemp)
	fmt.Println(eventT.QualMatches)
}
