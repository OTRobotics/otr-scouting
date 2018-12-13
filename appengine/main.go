package main

import (
	"github.com/gin-gonic/gin"
	"github.com/otrobotics/otr-scouting"
	"google.golang.org/appengine"
	"net/http"
)

type PageData struct {
	PageTitle string
}

type IndexPage struct {
	PageTitle string
	Events    []otrscouting.EventTemplate
}

func main() {
	router := gin.Default()
	router.GET("/", GinHomeHandler)
	router.GET("/upload", otrscouting.GinUploadHandler)
	router.POST("/upload", otrscouting.GinUploadHandler)
	router.GET("/upload/type", otrscouting.GinDataTypeHandler)

	router.GET("/matches", otrscouting.GinMatchesHandler)
	router.GET("/teams", otrscouting.GinTeamsHandler)
	router.GET("/events", otrscouting.GinEventsHandler)

	router.GET("/match/:matchnumber", otrscouting.GinMatchHandler)
	router.GET("/team/:teamnumber", otrscouting.GinTeamHandler)
	router.GET("/event/:event", otrscouting.GinEventHandler)

	router.POST("/slack/slash", otrscouting.GinSlackSlashHandler)

	http.Handle("/", router)
	appengine.Main()
}

func GinHomeHandler(c *gin.Context) {
	tmpl := otrscouting.GetPageTemplate("index.html", c)
	tmpl.Execute(c.Writer, nil)
}
