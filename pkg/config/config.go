package config

import (
	"errors"
	"os"

	env "github.com/joho/godotenv"
	"github.com/rwandaopensource/botx/pkg/helper"
)

// ErrEnv error returned when one the variables are missing
var ErrEnv error = errors.New("one or more of the environments are missing")

// Config load and validate missing variables
// if enforce is set to true it will exists the program when there is missing varibles
func Config(enforce bool) {
	var (
		mode string = os.Getenv("GO_ENV")
		file        = ".env"
	)
	// if you want to enforce the environment during tests don't run set GO_ENV=test
	enforce = mode == "test"
	switch mode {
	case "development", "production", "test":
		file += ("." + mode)
	default:
		break
	}
	if err := env.Load(); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			helper.Verbose("environment file not found; using os environment")
		} else {
			helper.FatalError(err, "")
		}
	}
	requiredEnv := []string{
		"DATABASE_URL",
		"DATABASE_NAME",
		"SLACK_CLIENT_ID",
		"SLACK_CLIENT_SECRET",
		"PRIVATE_KEY",
		"PUBLIC_KEY",
	}
	var missingEnv string
	for _, key := range requiredEnv {
		if os.Getenv(key) == "" {
			missingEnv = missingEnv + ", " + key
		}
	}
	if missingEnv != "" && enforce {
		helper.FatalError(ErrEnv, missingEnv)
	}
}
