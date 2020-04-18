package helper

import (
	"fmt"
	"log"
	"os"
)

// Print regardless of verbose flag
func Print(v ...interface{}) {
	d := fmt.Sprint(v...)
	log.Printf("\033[1;32m%s\033[0m\n", d)
}

// Verbose print everything when the verbose flag is set to true
func Verbose(v ...interface{}) {
	if v := os.Getenv("VERBOSE"); v != "" {
		Print(v)
	}
}

// FatalError print error(if any) and exists the program
func FatalError(err error, m string) {
	if err != nil {
		log.Fatalf("\033[1;31m%v \033[0m%s\n", err, m)
	}
}

// PrintError just print an error(if any) as warning
func PrintError(err error, m string) {
	if err != nil {
		log.Printf("\033[1;33m%v \033[0m%s\n", err, m)
	}
}
