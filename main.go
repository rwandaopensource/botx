package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/rwandaopensource/botx/database"
	"github.com/rwandaopensource/botx/helper"
	"github.com/rwandaopensource/botx/route"
	"github.com/rwandaopensource/botx/util"
)

func main() {

	defer database.Close()
	if util.Command() {
		return
	}

	ADDR := os.Getenv("PORT")
	if ADDR == "" {
		ADDR = ":8080"
	}
	s := &http.Server{
		Addr:           ADDR,
		Handler:        route.NewRouter(),
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	helper.Print("starting server", ADDR)

	go func() {
		helper.FatalError(s.ListenAndServe(), "")
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	helper.Verbose("gracefully shutting down")
	s.Shutdown(ctx)
	return
}
