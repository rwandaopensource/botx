package config

import (
	"errors"
	"log"
	"os"
	"strings"

	env "github.com/joho/godotenv"
)

// ErrEnv error returned when one the variables are missing
var ErrEnv error = errors.New("one or more of the environments are missing")

// Env contains all necessary environment variables of this application
var Env map[string]string = map[string]string{}

func init() {
	var (
		mode string = os.Getenv("GO_ENV")
		file        = ".env"
	)
	switch mode {
	case "development", "production", "test":
		file += ("." + mode)
	default:
		break
	}
	if err := env.Load(); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Println("environment file not found using os environment")
		} else {
			log.Fatalln(err)
		}
	}
	config()
}

// validate missing variables
func config() {
	Env["DATABASE_URL"] = os.Getenv("DATABASE_URL")
	Env["DATABASE_NAME"] = os.Getenv("DATABASE_NAME")
	var missingEnv string
	for key, value := range Env {
		if value == "" {
			missingEnv = missingEnv + ", " + key
		}
	}
	if missingEnv != "" {
		log.Fatalln(ErrEnv, strings.Trim(missingEnv, ", "))
	}
}
