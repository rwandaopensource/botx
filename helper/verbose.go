package helper

import (
	"fmt"
	"log"
	"os"
	"testing"
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

// PanicError it panics
func PanicError(err error, m string) {
	if err != nil {
		log.Panicf("\033[1;31m%v \033[0m%s\n", err, m)
	}
}

// PrintError just print an error(if any) as warning
func PrintError(err error, m string) {
	if err != nil {
		log.Printf("\033[1;33m%v \033[0m%s\n", err, m)
	}
}

//TestLog equivalent to t.Log with colors
func TestLog(t *testing.T, v ...interface{}) {
	d := fmt.Sprint(v...)
	t.Logf("\033[1;32m%s\033[0m\n", d)
}

// TestLogf equivalent to t.Logf with colors
func TestLogf(t *testing.T, format string, v ...interface{}) {
	format = "\033[1;32m" + format + "033[0m\n"
	t.Logf(format, v...)
}

// TestError equivalent to t.Error with colors
func TestError(t *testing.T, v ...interface{}) {
	d := fmt.Sprint(v...)
	t.Errorf("\033[1;31m%s\033[0m\n", d)
}

// TestErrorf equivalent to t.Errorf with colors
func TestErrorf(t *testing.T, format string, v ...interface{}) {
	format = "\033[1;31m" + format + "033[0m\n"
	t.Errorf(format, v...)
}
