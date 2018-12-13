package otrscouting

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type EventTemplate struct {
	EventName   string
	FRCEvents   string
	EventCode   string
	EventDate   string
	Year        int
	QualMatches []MatchTemplate
	TeamTable   []PowerUpRobot
}

type MatchTemplate struct {
	MatchNumber string
	Red1        RobotTemplate
	Red2        RobotTemplate
	Red3        RobotTemplate
	Blue1       RobotTemplate
	Blue2       RobotTemplate
	Blue3       RobotTemplate
	redScore    int // Unused ATM
	blueScore   int // Unused ATM
	MatchId     string
}

func (m MatchTemplate) Level() string {
	if strings.Split(m.MatchId, "_")[2][0] == 'q' {
		return "Qualifications"
	}
	return "Playoffs"
}

func (m MatchTemplate) FriendlyName() string {
	return m.Level() + ": Match " + m.MatchNumber
}

func (m MatchTemplate) EventId() string {
	return strings.Split(m.MatchId, "_")[1]
}

func (m MatchTemplate) Year() int {
	year, _ := strconv.Atoi(strings.Split(m.MatchId, "_")[0])
	return year
}

// Format: /event/:event
func GinEventHandler(c *gin.Context) {
	eventCode := c.Param("event")
	tmpl := GetPageTemplate("event.html", c)
	data := getEvent(c, eventCode)
	data.QualMatches = getEventMatches(c, eventCode)
	data.TeamTable = sumeventrobotsPowerup(c, getEventRobots(c, eventCode))
	tmpl.Execute(c.Writer, data)
}
