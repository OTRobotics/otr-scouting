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
	router.GET("/match/:matchnumber", otrscouting.GinMatchHandler)
	router.GET("/team/:teamnumber", otrscouting.GinTeamHandler)
	router.GET("/event/:event", otrscouting.GinEventHandler)

	http.Handle("/", router)
	appengine.Main() // Starts the server to receive requests
}

func GinHomeHandler(c *gin.Context) {
	tmpl := otrscouting.GetPageTemplate("index.html", c)
	data := IndexPage{PageTitle: "OTR Scouting Application",
		Events: []otrscouting.EventTemplate{
			{EventName: "Waterloo District", EventCode: "2018_onwat", FRCEvents: "ONWAT", EventDate: "Week 4"},
			{EventName: "McMaster District", EventCode: "2018_onham", FRCEvents: "ONHAM", EventDate: "Week 6"},
		},
	}
	tmpl.Execute(c.Writer, data)
}
