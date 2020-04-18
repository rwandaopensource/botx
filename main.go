package main

import (
	"context"
	"flag"
	"os"
	"strings"

	"github.com/rwandaopensource/botx/pkg/database"
	"github.com/rwandaopensource/botx/pkg/helper"
)

func main() {
	flag.BoolVar(&drop, "drop-all", false, "drop all tables, should be done before running tests")
	flag.StringVar(&dropT, "drop", "", "drop tables, defined after this flag (comma separated)")
	flag.BoolVar(&drop, "d", false, "(shorthand) drop")
	flag.BoolVar(&help, "help", false, "dispaly information on commands")
	flag.BoolVar(&help, "h", false, "(shorthand) help")
	flag.BoolVar(&verbose, "verbose", false, "run the program in verbose mode")
	flag.BoolVar(&verbose, "v", false, "(shorthand) verbose")
	flag.Parse()
	if verbose {
		os.Setenv("VERBOSE", "true")
	}
	if command() {
		return
	}
	defer closeOpenHandle()
}

func command() bool {
	cmd := false
	database.InitDB()
	if drop {
		err := database.Drop()
		helper.PrintError(err, "")
		cmd = true
	} else if dropT != "" {
		err := database.DropSome(strings.Split(strings.Replace(dropT, " ", "", -1), ","))
		helper.PrintError(err, "")
		cmd = true
	}
	if help {
		helper.Print(usage)
		cmd = true
	}
	return cmd
}

func closeOpenHandle() {
	err := database.Client.Disconnect(context.TODO())
	helper.FatalError(err, "")
}

type table string

var (
	drop  bool
	dropT string
	usage string = `
Usage: botx [options]

Options:
-h, --help: display this usage
--drop-all: drop all tables, should be done before running tests,
            if parsed with --drop, only --drop-all will be executed
-d, --drop: drop tables, defined after this flag (comma separated)
            if parsed with --drop-all, only --drop-all only will be executed
-v, --verbose: run the program in verbose mode
	`
	help    bool
	verbose bool
)
