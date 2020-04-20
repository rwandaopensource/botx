package util

import (
	"flag"
	"strings"

	"github.com/rwandaopensource/botx/database"
	"github.com/rwandaopensource/botx/helper"
)

// Command parse and run command line command
func Command() bool {
	flag.BoolVar(&drop, "drop-all", false, "drop all tables, should be done before running tests")
	flag.StringVar(&dropT, "drop", "", "drop tables, defined after this flag (comma separated)")
	flag.BoolVar(&drop, "d", false, "(shorthand) drop")
	flag.BoolVar(&help, "help", false, "dispaly information on commands")
	flag.BoolVar(&help, "h", false, "(shorthand) help")
	flag.BoolVar(&verbose, "verbose", false, "run the program in verbose mode")
	flag.BoolVar(&verbose, "v", false, "(shorthand) verbose")
	flag.BoolVar(&key, "key", false, "generate new public and private key")
	flag.BoolVar(&key, "k", false, "(shorthand) key")
	flag.Parse()
	return command()
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
	if key {
		helper.PrintKey()
		cmd = true
	}
	return cmd
}

var (
	drop    bool
	dropT   string
	help    bool
	verbose bool
	key     bool
	usage   string = `
Usage: botx [options]

Options:
-h, --help: display this usage
--drop-all: drop all tables, should be done before running tests,
            if parsed with --drop, only --drop-all will be executed
-d, --drop: drop tables, defined after this flag (comma separated)
            if parsed with --drop-all, only --drop-all will be executed
-v, --verbose: run the program in verbose mode
-k, --key: generate new public and private key
	`
)
