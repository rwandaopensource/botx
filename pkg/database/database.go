package database

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/rwandaopensource/botx/pkg/config"
	"github.com/rwandaopensource/botx/pkg/helper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client is the mongodb client of the whole application
var Client *mongo.Client

// DB default database to use
var DB *mongo.Database

// InitDB initiate the database connection and perform a warmup connection
func InitDB() {
	config.Config(true)
	helper.Verbose("all env variables ✅!")
	var err error
	var (
		errDBOptions = errors.New("invalid database options")
		errDBPing    = errors.New("could not ping the db connection")
	)
	dbURL := os.Getenv("DATABASE_URL")
	dbName := os.Getenv("DATABASE_NAME")
	opts := options.Client()
	opts.ApplyURI(dbURL)
	opts.SetMaxPoolSize(0)
	opts.SetRetryReads(true)
	opts.SetRetryWrites(true)
	opts.SetSocketTimeout(time.Second * 5)
	opts.SetMaxConnIdleTime(time.Second * 5)
	if opts.Validate() != nil {
		helper.FatalError(errDBOptions, "")
	}
	Client, err = mongo.Connect(context.TODO(), opts)
	helper.FatalError(err, "")
	helper.Verbose("db connection warm-up")
	if Client.Ping(context.TODO(), nil) != nil {
		helper.FatalError(errDBPing, "")
	}
	if DB = Client.Database(dbName); DB != nil {
		helper.Verbose("db connected ✅")
	}
}
