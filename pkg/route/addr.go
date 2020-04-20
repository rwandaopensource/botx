package route

import (
	"os"
)

// ADDR of the server
var ADDR string

func init() {
	addr := os.Getenv("PORT")
	if addr == "" {
		ADDR = ":8080"
		return
	}
	ADDR = ":" + addr
}
