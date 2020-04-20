package slack

import (
	"fmt"
	"os"
)

// InstallHTML return html for installing app
func InstallHTML() string {
	slackClientID := os.Getenv("SLACK_CLIENT_ID")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	host := "http://localhost:" + port
	return fmt.Sprintf(`
<html lang="en">
<head>
<title> Install slack </tile>
</head>
<h1> Install this app in your slack workspace </h1>

<a href="https://slack.com/oauth/authorize?client_id=%s&scope=im:write,chat:write,users:read,reactions:read&redirect_uri=%s"
</html>	
`, slackClientID)
}
