package main

import (
	"github.com/rwandaopensource/botx/pkg/database"
	"github.com/rwandaopensource/botx/pkg/util"
)

func main() {

	defer database.CloseDB()
	if util.Command() {
		return
	}
}
