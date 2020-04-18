package database

import (
	"context"

	"github.com/rwandaopensource/botx/pkg/helper"
)

// Drop will drop all tables, suitable before running tests
func Drop() error {
	for _, t := range Tables {
		if err := DB.Collection(t).Drop(context.TODO()); err != nil {
			return err
		}
		helper.Verbose("dropped " + t + " table")
	}
	return nil
}

// DropSome drops tables that are parsed in t params
func DropSome(t []string) error {
	for _, v := range t {
		if err := DB.Collection(v).Drop(context.TODO()); err != nil {
			return err
		}
		helper.Verbose("dropped " + v + " table")
	}
	return nil
}
