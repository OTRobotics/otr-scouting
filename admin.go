package otrscouting

import "github.com/gin-gonic/gin"

type AdminTemplate struct {
	PageTitle string
}

func GoAdminHandler(c *gin.Context) {
	// Auth flow:
	/*
		If no auth - message to say need auth - Use /scoutingadmin in #otr-scouting.
		Slack slash command sends req. here - create jwt for one time use, and respond to Slack.
		Follow: https://github.com/appleboy/gin-jwt
		Valid auth permit for 5m.
	*/
	var data AdminTemplate
	data.PageTitle = "Admin Page"
	tmpl := GetPageTemplate("admin.html", c)

	tmpl.Execute(c.Writer, data)

	// Delete Station Data
	// Overwrite Station Data
}
