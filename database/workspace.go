package database

import (
	"context"

	"github.com/slack-go/slack"
)

// Workspace represent workspace
type Workspace struct {
	slack.OAuthResponse
	CreatedAt int `json:"created_at"`
	UpdatedAt int `json:"updated_at"`
}

// CreateWorkspace store workspace details in the database
func CreateWorkspace(workspace Workspace) error {
	_, err := WorkspaceModel.InsertOne(context.TODO(), workspace, nil)
	if err != nil {
		return err
	}
	return nil
}
