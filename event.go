package otrscouting

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
	Red1      RobotTemplate
	Red2      RobotTemplate
	Red3      RobotTemplate
	Blue1     RobotTemplate
	Blue2     RobotTemplate
	Blue3     RobotTemplate
	RedScore  int
	BlueScore int
}

type RobotTemplate struct {
}

// Format: /event/:event
func GinEventHandler(c *gin.Context) {
	event := c.Param("event")
	c.String(http.StatusOK, event)
}
