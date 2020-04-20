package helper

import (
	"fmt"
	"os"
)

// InstallHTML return html for installing app
func InstallHTML() string {
	slackClientID := os.Getenv("SLACK_CLIENT_ID")
	host := os.Getenv("REDIRECT_URI")
	return fmt.Sprintf(`
<html lang="en">
<head>
<title> Install slack </tile>
</head>
<h1> Install this app in your slack workspace </h1>
<a href="https://slack.com/oauth/v2/authorize?client_id=%s&scope=im:write,chat:write,users:read,reactions:read,chat:write.public,incoming-webhook&redirect_uri=%s">
<img alt="Add to Slack" height="40" width="139" src="https://platform.slack-edge.com/img/add_to_slack.png" srcset="https://platform.slack-edge.com/img/add_to_slack.png 1x, https://platform.slack-edge.com/img/add_to_slack@2x.png 2x">
</a>
</html>
`, slackClientID, host)
}
