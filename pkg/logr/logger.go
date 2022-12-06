package logr

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/exp/slog"
)

var (
	log         *slog.Logger
	logRandomID = uuid.New().String()
)

func init() {
	log = slog.Default()
}

func Default() *slog.Logger {
	return log
}

func WithContext(ctx context.Context, log *slog.Logger) context.Context {
	gctx, ok := ctx.(*gin.Context)
	if ok {
		gctx.Set(logRandomID, log)
		return gctx
	}

	return context.WithValue(ctx, logRandomID, log)
}

func FromContext(ctx context.Context) *slog.Logger {
	log := ctx.Value(logRandomID)
	logr, ok := log.(*slog.Logger)
	if ok {
		return logr
	}

	return Default()
}
