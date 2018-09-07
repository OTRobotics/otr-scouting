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
	ElimMatches []MatchTemplate
}

type MatchTemplate struct {
	MatchNumber string
	Red1        RobotTemplate
	Red2        RobotTemplate
	Red3        RobotTemplate
	Blue1       RobotTemplate
	Blue2       RobotTemplate
	Blue3       RobotTemplate
	RedScore    int
	BlueScore   int
	MatchId     string
}

func (m MatchTemplate) Level() string {
	if strings.Split(m.MatchId, "_")[2][0] == 'q' {
		return "Qualification"
	}
	return "Elimination"
}

func (m MatchTemplate) FriendlyName() string {
	return m.Level() + " " + m.MatchNumber
}

func (m MatchTemplate) EventId() string {
	return strings.Split(m.MatchId, "_")[1]
}

func (m MatchTemplate) Year() int {
	year, _ := strconv.Atoi(strings.Split(m.MatchId, "_")[0])
	return year
}

type RobotTemplate struct {
	Team int `json:"team"`
}

// Format: /event/:event
func GinEventHandler(c *gin.Context) {
	eventCode := c.Param("event")
	tmpl := GetPageTemplate("event.html", c)
	data := getEvent(c, eventCode)
	data.QualMatches = getEventMatches(c, eventCode)
	tmpl.Execute(c.Writer, data)
}
