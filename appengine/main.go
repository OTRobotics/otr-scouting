package main

import (
	"github.com/gin-gonic/gin"
	"github.com/otrobotics/otr-scouting"
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

	router.GET("/matches")
	router.GET("/teams")
	router.GET("/events")

	router.GET("/match/:matchnumber", otrscouting.GinMatchHandler)
	router.GET("/team/:teamnumber", otrscouting.GinTeamHandler)
	router.GET("/event/:event", otrscouting.GinEventHandler)

	http.Handle("/", router)
}

func GinHomeHandler(c *gin.Context) {
	tmpl := otrscouting.GetPageTemplate("index.html", c)
	data := IndexPage{PageTitle: "OTR Scouting Application",
		Events: otrscouting.GetYearEvents(2018, c),
	}
	tmpl.Execute(c.Writer, data)
}
