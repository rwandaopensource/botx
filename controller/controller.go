package controller

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/rwandaopensource/botx/database"
	"github.com/rwandaopensource/botx/helper"
	"github.com/slack-go/slack"
)

var (
	slackSigningSecret = os.Getenv("SLACK_SIGNING_SECRET")
	slackClientID      = os.Getenv("SLACK_CLIENT_ID")
	slackClientSecret  = os.Getenv("SLACK_CLIENT_SECRET")
	slackError         = "invalid_client_id,bad_client_secret,oauth_authorization_url_mismatch,bad_redirect_uri"
)

// Install http.handler for installer command
func Install(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	code := r.URL.Query().Get("code")
	e := r.URL.Query().Get("error")
	if e != "" {
		if strings.Contains(slackError, e) {
			rw.WriteHeader(http.StatusInternalServerError)
			helper.PrintError(errors.New(e), "during install")
			return
		}
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(rw, strings.Join(strings.Split(e, "_"), " "))
		return
	}
	t, err := slack.GetOAuthResponse(http.DefaultClient, slackClientID, slackClientSecret, code, "")
	if err != nil {
		rw.WriteHeader(http.StatusNotAcceptable)
		return
	}
	team := database.Workspace{*t, time.Now().Nanosecond(), time.Now().Nanosecond()}
	err = database.CreateWorkspace(team)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
	}
}
