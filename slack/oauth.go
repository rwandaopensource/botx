package slack

import (
	"net/http"
	"os"

	"github.com/slack-go/slack"
)

var (
	slackClientID     = os.Getenv("SLACK_CLIENT_ID")
	slackClientSecret = os.Getenv("SLACK_CLIENT_SECRET")
)

// OAuthToken return slack authentication on a worskpace
func OAuthToken(code string, redirectURI string) (string, string, error) {
	return slack.GetOAuthToken(http.DefaultClient, slackClientID, slackClientSecret, code, redirectURI)
}
