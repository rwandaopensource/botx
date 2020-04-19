package database

import (
	"context"

	"github.com/rwandaopensource/botx/pkg/helper"
)

// CloseDB releases connection open by database
func CloseDB() {
	err := Client.Disconnect(context.TODO())
	helper.FatalError(err, "")
}
