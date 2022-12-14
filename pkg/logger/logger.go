package logger

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-jarvis/logr"
)

var loggerKey = "rum-gonic.logger.key"

func FromContext(ctx context.Context) logr.Logger {
	val := ctx.Value(loggerKey)
	log, ok := val.(logr.Logger)
	if ok {
		return log
	}

	return logr.FromContext(ctx)
}

func WithLogger(ctx context.Context, log logr.Logger) context.Context {
	ginc, ok := ctx.(*gin.Context)
	if ok {
		ginc.Set(loggerKey, log)
		return ginc
	}

	return context.WithValue(ctx, loggerKey, log)
}
