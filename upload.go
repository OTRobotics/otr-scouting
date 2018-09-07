package otrscouting

import (
	"cloud.google.com/go/datastore"
	"context"
	json2 "encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type UploadData struct {
	PageTitle    string
	Uploaded     bool
	Unauthorized bool
}

type MatchUpload struct {
	MatchId   string `json:"match_id"`
	Level     string `json:"comp_level"`
	EventId   string `json:"event_id"`
	Year      int    `json:"year"`
	MatchData struct {
		Red1  RobotTemplate `json:"red_1"`
		Red2  RobotTemplate `json:"red_2"`
		Red3  RobotTemplate `json:"red_3"`
		Blue1 RobotTemplate `json:"blue_1"`
		Blue2 RobotTemplate `json:"blue_2"`
		Blue3 RobotTemplate `json:"blue_3"`
	} `json:"match_data"`
}

func GinUploadHandler(c *gin.Context) {
	tmpl := GetPageTemplate("upload.html", c)
	data := UploadData{
		PageTitle:    "Upload Event Info",
		Uploaded:     false,
		Unauthorized: true,
	}
	if c.Request.Method == "GET" {
		tmpl.Execute(c.Writer, data)
	} else if c.Request.Method == "POST" {
		data.Uploaded = true
		// Extract JSON
		pass := c.PostForm("pass")
		if pass == "OTR3474" {
			data.Unauthorized = false
		}
		if data.Unauthorized {
			tmpl.Execute(c.Writer, data)
			return
		}
		rawJson := c.PostForm("json")
		var matches []MatchUpload
		tmpl.Execute(c.Writer, data)
		if err := json2.Unmarshal([]byte(rawJson), &matches); err != nil {
			fmt.Fprint(c.Writer, err)
		}

		uploadMatchToDatastore(c, matches)

	}
}

func uploadMatchToDatastore(c *gin.Context, matches []MatchUpload) {
	// Set your Google Cloud Platform project ID.
	projectID := "otr-scouting"
	ctx := context.Background()
	// Creates a client.
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Set type
	kind := "match"
	for _, match := range matches {
		// Fully qualified match id - {year}_{event}_{matchid}
		name := strconv.Itoa(match.Year) + "_" + match.EventId + "_" + match.MatchId
		// Creates a Key instance.
		taskKey := datastore.NameKey(kind, name, nil)

		// Saves the new entity.
		if _, err := client.Put(c, taskKey, &match); err != nil {
			fmt.Fprintf(c.Writer, "Failed to save match: %v", err)
		}

		fmt.Fprintf(c.Writer, "Saved %v: %v\n", taskKey, name)
	}
}
