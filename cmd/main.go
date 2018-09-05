package cmd

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"net/http"
	"otr-scouting"
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
	data := PageData{PageTitle: "OTR Scouting Application"}
	tmpl.Execute(c.Writer, data)
}
