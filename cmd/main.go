package cmd

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"net/http"
)

type PageData struct {
	PageTitle string
}

type IndexPage struct {
	PageTitle string
	Events    []EventTemplate
}

func main() {
	router := gin.Default()
	router.GET("/", GinHomeHandler)
	router.GET("/upload", GinUploadHandler)
	router.POST("/upload", GinUploadHandler)
	router.GET("/match/:matchnumber", GinMatchHandler)
	router.GET("/team/:teamnumber", GinTeamHandler)
	router.GET("/event/:event", GinEventHandler)

	http.Handle("/", router)
	appengine.Main() // Starts the server to receive requests
}

func GinHomeHandler(c *gin.Context) {
	tmpl := GetPageTemplate("index.html", c)
	data := PageData{PageTitle: "OTR Scouting Application"}
	tmpl.Execute(c.Writer, data)
}
