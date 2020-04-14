package main

import (
	"context"
	"log"

	"github.com/rwandaopensource/botx/pkg/database"
)

func main() {
	defer closeOpenHandle()
}

func closeOpenHandle() {
	err := database.Client.Disconnect(context.TODO())
	if err != nil {
		log.Fatalln(err)
	}
}
