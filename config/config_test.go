package config_test

import (
	"os"
	"testing"

	"github.com/rwandaopensource/botx/config"
	"github.com/rwandaopensource/botx/helper"
)

func TestConfigParseError(t *testing.T) {
	tmp := os.Getenv("PRIVATE_KEY")
	os.Setenv("PRIVATE_KEY", "p")
	config.Config(false)
	if v := os.Getenv("PRIVATE_KEY"); v != "p" {
		helper.TestError(t, "parsing config failed")
	}
	os.Setenv("PRIVATE_KEY", tmp)
}
