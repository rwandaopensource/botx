package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/rwandaopensource/botx/pkg/database"
	"github.com/rwandaopensource/botx/pkg/helper"
	"github.com/rwandaopensource/botx/pkg/route"
	"github.com/rwandaopensource/botx/pkg/util"
)

func main() {

	defer database.CloseDB()
	if util.Command() {
		return
	}

	s := &http.Server{
		Addr:           route.ADDR,
		Handler:        route.Router(),
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	helper.Print("starting server", route.ADDR)

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
