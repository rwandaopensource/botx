package database

import (
	"errors"

	"github.com/rwandaopensource/botx/pkg/helper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// User represent a user
type User struct {
	// ClientID is the id signed to the user
	ClientID string
	// ClientSecret is the secret signed to the user
	ClientSecret string
	// TeamID slack workspace id

}

// ReadUser read a single user
func ReadUser(clientID string) (*User, error) {
	ctx, cancel := helper.ReadRecordCtx()
	defer cancel()
	res := UserModel.FindOne(ctx, bson.M{"clientId": clientID}, nil)
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
