package helper

import (
	"context"
	"time"
)

// ReadRecordCtx is the context that set the read timeout
func ReadRecordCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Second*5)
}

// WriteRecordCtx is the context that set the write timeout
func WriteRecordCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Second*5)
}
