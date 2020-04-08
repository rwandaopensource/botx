package database

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/rwandaopensource/botx/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client is the mongodb client of the whole application
var Client *mongo.Client

// DB default database to use
var DB *mongo.Database

func init() {
	log.Println("all env variables ✅!")
	var err error
	var (
		errDBOptions = errors.New("invalid database options")
		errDBPing    = errors.New("could not ping the db connection")
	)
	dbURL := config.Env["DATABASE_URL"]
	dbName := config.Env["DATABASE_NAME"]
	opts := options.Client()
	opts.ApplyURI(dbURL)
	opts.SetMaxPoolSize(0)
	opts.SetRetryReads(true)
	opts.SetRetryWrites(true)
	opts.SetSocketTimeout(time.Second * 5)
	opts.SetMaxConnIdleTime(time.Second * 5)
	if opts.Validate() != nil {
		log.Fatalln(errDBOptions)
	}
	Client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatalln("db connection error:", err)
	}
	err = Client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalln(errDBPing, err)
	}
	DB = Client.Database(dbName)
	if DB != nil {
		log.Println("db connected ✅!")
	}
}
