package otrscouting

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

// FORMAT: /match/:matchnumber
func GinMatchHandler(c *gin.Context) {
	name := c.Param("matchnumber")
	// Check match number/format against event/match
	// 2018_event_match
	split := strings.Split(name, "_")
	if len(split) < 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Match Number Format"})
	}
	year := split[0]
	event := split[1]
	match := split[2]
	c.JSON(http.StatusOK, gin.H{"match": match, "event": event, "year": year})
}

func getMatchData(year int, event string, match string) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "otr-scouting")
	if err != nil {
		// TODO: Handle error.
	}

	eventRef := client.Collection(strconv.Itoa(year) + "_" + event)
	matchRef := eventRef.Doc(match)
	docsnap, err := matchRef.Get(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	dataMap := docsnap.Data()
	fmt.Println(dataMap)
}
