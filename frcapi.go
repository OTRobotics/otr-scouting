package otrscouting

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
	"io/ioutil"
	"net/http"
	"strings"
)

func getMatchScore(match string, c *gin.Context) {
	makeTBARequest(c, "/match/"+strings.Replace(match, "_", "", -1))
}

func makeTBARequest(c *gin.Context, path string) string {
	url := "https://www.thebluealliance.com/api/v3" + path
	ctx := appengine.NewContext(c.Request)
	client := urlfetch.Client(ctx)
	req, err := http.NewRequest("GET", url, nil)

	req.Header.Set("X-TBA-Auth-Key", "")
	resp, err := client.Do(req)
	if err != nil {
		return ""
	}

	body, _ := ioutil.ReadAll(resp.Body)
	log.Debugf(ctx, "HTTP GET returned status %v", resp.Status)
	return string(body)
}
