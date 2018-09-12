package otrscouting

import (
	"cloud.google.com/go/datastore"
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"html/template"
	"io/ioutil"
	"strconv"
	"strings"
)

func readFile(fileName string, c *gin.Context) []byte {
	ctx := appengine.NewContext(c.Request)
	bucketName := "staging.otr-scouting.appspot.com"
	//[END get_default_bucket]
	creds, err := google.FindDefaultCredentials(ctx, storage.ScopeReadOnly)
	if err != nil {
		log.Errorf(ctx, "%p", err)
	}
	client, err := storage.NewClient(ctx, option.WithCredentials(creds))
	bucket := client.Bucket(bucketName)
	rc, err := bucket.Object(fileName).NewReader(ctx)
	if err != nil {
		//errorf("readFile: unable to open file from bucket %q, file %q: %v", d.bucketName, fileName, err)
		return []byte("")
	}
	defer rc.Close()
	slurp, err := ioutil.ReadAll(rc)
	if err != nil {
		//d.errorf("readFile: unable to read data from bucket %q, file %q: %v", d.bucketName, fileName, err)
		return []byte("")
	}

	return slurp
}

func GetPageTemplate(page string, c *gin.Context) *template.Template {

	page = "web/" + page
	// reads html as a slice of bytes
	html := readFile(page, c)

	tmpl, _ := template.New(page).Parse(string(html))
	return tmpl
}

func datastoreClient(c *gin.Context) *datastore.Client {
	projectID := "otr-scouting"
	ctx := context.Background()
	// Creates a client.
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Errorf(ctx, "Failed to create client: %v", err)
	}
	return client
}

func uploadMatchToDatastore(c *gin.Context, matches []MatchUpload) {
	// Set your Google Cloud Platform project ID.

	client := datastoreClient(c)
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

		fmt.Fprintf(c.Writer, "Saved %v: %v<br>", taskKey, name)

		var robots []RobotTemplate
		robots = append(robots, match.MatchData.Red1)
		robots = append(robots, match.MatchData.Red2)
		robots = append(robots, match.MatchData.Red3)
		robots = append(robots, match.MatchData.Blue1)
		robots = append(robots, match.MatchData.Blue2)
		robots = append(robots, match.MatchData.Blue3)

		uploadRobotToDatastore(c, robots, name)
	}
}

func uploadRobotToDatastore(c *gin.Context, robots []RobotTemplate, matchcode string) {
	client := datastoreClient(c)
	// Set type
	kind := "robot"
	for _, robot := range robots {
		// Fully qualified match/team - {year}_{event}_{matchid}_{team}
		name := matchcode + "_" + strconv.Itoa(robot.Team)
		// Creates a Key instance.
		taskKey := datastore.NameKey(kind, name, nil)

		// Saves the new entity.
		if _, err := client.Put(c, taskKey, &robot); err != nil {
			fmt.Fprintf(c.Writer, "Failed to save robot: %v", err)
		}

		fmt.Fprintf(c.Writer, "Saved %v: %v<br>", taskKey, name)
	}
}

func getEvent(c *gin.Context, event string) EventTemplate {
	ctx := context.Background()
	var eTemp EventTemplate

	q := datastore.NewQuery("event").
		Filter("EventCode =", event)
	client := datastoreClient(c)
	keys, err := client.GetAll(ctx, q, &eTemp)

	if err != nil {
		log.Errorf(appengine.NewContext(c.Request), "datastoredb: could not list books: %v", err)
	}

	log.Debugf(appengine.NewContext(c.Request), "Found Keys: %s", keys)

	eTemp.QualMatches = getEventQualMatches(c, event)
	eTemp.ElimMatches = getEventElimMatches(c, event)

	return eTemp
}

func GetYearEvents(year int, c *gin.Context) []EventTemplate {
	ctx := context.Background()
	eTemp := make([]EventTemplate, 0)

	q := datastore.NewQuery("event").
		Filter("Year =", year)
	client := datastoreClient(c)
	keys, err := client.GetAll(ctx, q, &eTemp)

	if err != nil {
		log.Errorf(appengine.NewContext(c.Request), "datastoredb: could not list: %v", err)
	}

	log.Debugf(appengine.NewContext(c.Request), "Found Keys: %s", keys)

	return eTemp
}

func getEventQualMatches(c *gin.Context, event string) []MatchTemplate {
	split := strings.Split(event, "_")

	eventId := split[1]
	year, _ := strconv.Atoi(split[0])

	ctx := context.Background()
	matches := make([]*MatchUpload, 0)
	q := datastore.NewQuery("match").
		Filter("Year =", year).
		Filter("EventId =", eventId).
		Filter("Level =", "qualifications")

	client := datastoreClient(c)
	keys, err := client.GetAll(ctx, q, &matches)

	if err != nil {
		log.Errorf(appengine.NewContext(c.Request), "datastoredb: could not list: %v", err)
	}

	log.Debugf(appengine.NewContext(c.Request), "Found Keys: %s", keys)
	var cleanedMatches []MatchTemplate

	for _, v := range matches {
		cleanedMatches = append(cleanedMatches, v.toMatchTemplate())
	}
	return cleanedMatches
}

func getEventElimMatches(c *gin.Context, event string) []MatchTemplate {
	split := strings.Split(event, "_")

	eventId := split[1]
	year, _ := strconv.Atoi(split[0])

	ctx := context.Background()
	matches := make([]*MatchUpload, 0)
	q := datastore.NewQuery("match").
		Filter("Year =", year).
		Filter("EventId =", eventId).
		Filter("Level =", "eliminations")

	client := datastoreClient(c)
	keys, err := client.GetAll(ctx, q, &matches)

	if err != nil {
		log.Errorf(appengine.NewContext(c.Request), "datastoredb: could not list: %v", err)
	}

	log.Debugf(appengine.NewContext(c.Request), "Found Keys: %s", keys)
	var cleanedMatches []MatchTemplate

	for _, v := range matches {
		cleanedMatches = append(cleanedMatches, v.toMatchTemplate())
	}
	return cleanedMatches
}

func getEventMatches(c *gin.Context, event string) []MatchTemplate {
	split := strings.Split(event, "_")

	eventId := split[1]
	year, _ := strconv.Atoi(split[0])

	ctx := context.Background()
	matches := make([]*MatchUpload, 0)
	q := datastore.NewQuery("match").
		Filter("Year =", year).
		Filter("EventId =", eventId)
	client := datastoreClient(c)
	keys, err := client.GetAll(ctx, q, &matches)

	if err != nil {
		log.Errorf(appengine.NewContext(c.Request), "datastoredb: could not list: %v", err)
	}

	log.Debugf(appengine.NewContext(c.Request), "Found Keys: %s", keys)
	var cleanedMatches []MatchTemplate

	for _, v := range matches {
		cleanedMatches = append(cleanedMatches, v.toMatchTemplate())
	}
	return cleanedMatches
}

func getMatch(c *gin.Context, matchCode string) MatchTemplate {
	split := strings.Split(matchCode, "_")

	eventId := split[1]
	year, _ := strconv.Atoi(split[0])
	matchId := split[2][1:]
	level := split[2][0]
	compLevel := ""
	if level != 'e' {
		compLevel = "qualifications"
	} else {
		compLevel = "eliminations"
	}

	ctx := context.Background()
	matches := make([]MatchUpload, 0)
	q := datastore.NewQuery("match").
		Filter("Year =", year).
		Filter("EventId =", eventId).
		Filter("Level =", compLevel).
		Filter("MatchId = ", matchId)
	client := datastoreClient(c)
	keys, err := client.GetAll(ctx, q, &matches)

	if err != nil {
		log.Errorf(appengine.NewContext(c.Request), "datastoredb: could not list: %v", err)
	}

	log.Debugf(appengine.NewContext(c.Request), "Found Keys: %s", keys)
	if len(matches) == 0 {
		return MatchTemplate{}
	} else {
		return matches[0].toMatchTemplate()
	}
}

func sumPowerUpRobots(c *gin.Context, eventCode string) {

}
