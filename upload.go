package otrscouting

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GinUploadHandler(c *gin.Context) {
	if c.Request.Method == "GET" {
		tmpl := GetPageTemplate("upload.html", c)
		tmpl.Execute(c.Writer, nil)
	} else {
		// POST Request, serve JSON reply.
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	}
}
