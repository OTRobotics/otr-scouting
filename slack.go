package otrscouting

import (
	"github.com/gin-gonic/gin"
	"github.com/nlopes/slack"
)

var (
	slackClient *slack.Client = nil
)

type SlackCommandData struct {
	Channel     string
	Command     string
	User        string
	ResponseUrl string
	Args        string
}

func slackInit(c *gin.Context) {
	if slackClient == nil {
		botToken := ""
		slackClient = slack.New(botToken)
	}
}

func GinSlackSlashHandler(c *gin.Context) {
	slackInit(c)

	commandData := SlackCommandData{
		Channel:     c.PostForm("channel_id"),
		Command:     c.PostForm("command"),
		User:        c.PostForm("user_id"),
		ResponseUrl: c.PostForm("response_url"),
		Args:        c.PostForm("text"),
	}

	command := c.PostForm("command")
	switch command {
	case "/match":
		break
	case "/team":
		break
	case "/event":
		slashEvent(c, commandData)
		break
	default:
		break
	}
}

func slashEvent(c *gin.Context, data SlackCommandData) {
	if slackClient == nil {
		// Fail. No slack client to reply with other queries as needed.
		c.JSON(200, gin.H{"text": "https://otr-scouting.appspot.com/event/" + data.Args, "in_channel": true})
		return
	}
	attachment := slack.Attachment{
		Text:    "https://otr-scouting.appspot.com/event/" + data.Args,
	}
	_, _, err :=slackClient.PostMessage(data.Channel, slack.MsgOptionText("Event:", false), slack.MsgOptionAttachments(attachment))
	if err != nil {
		print(err)
		c.JSON(200, gin.H{"text": "https://otr-scouting.appspot.com/event/" + data.Args, "in_channel": true})
	}
}
