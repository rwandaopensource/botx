package database

import (
	"context"
	"errors"
	"time"

	"github.com/rwandaopensource/botx/helper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// User represent a user
type User struct {
	CreatedAt    time.Duration `json:"created_at"`
	UpdatedAt    time.Duration `json:"updated_at"`
	ClientID     string        `json:"client_id"`
	ClientSecret string        `json:"client_secret"`
	TeamID       string        `json:"team_id"`
}

// ReadUser read a single user
func ReadUser(clientID string) (*User, error) {
	res := UserModel.FindOne(context.TODO(), bson.M{"clientId": clientID}, nil)
	err := res.Err()
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}
	if err != nil {
		helper.PrintError(err, "Reading user")
		return nil, err
	}
	user := &User{}
	err = res.Decode(user)
	if err != nil {

	}
	return user, nil
}
