package otrscouting

import (
	"github.com/gin-gonic/gin"
)

type EventTemplate struct {
	EventName   string
	FRCEvents   string
	EventCode   string
	EventDate   string
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

type RobotTemplate struct {
}

// Format: /event/:event
func GinEventHandler(c *gin.Context) {
	//	event := c.Param("event")

	tmpl := GetPageTemplate("event.html", c)
	data := EventTemplate{
		EventName: "Waterloo District",
		FRCEvents: "ONWAT",
		EventCode: "2018_onwat",
		EventDate: "Week 4",
		QualMatches: []MatchTemplate{
			{
				MatchNumber: "1",
				RedScore:    120,
				BlueScore:   200,
				MatchId:     "2018_onwat_q1",
			},
		},
		ElimMatches: []MatchTemplate{
			{
				MatchNumber: "QF1-1",
				RedScore:    100,
				BlueScore:   90,
				MatchId:     "2018_onwat_qf1-1",
			},
		},
	}
	tmpl.Execute(c.Writer, data)
}
