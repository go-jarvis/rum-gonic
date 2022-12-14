package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/go-jarvis/logr"
	"github.com/go-jarvis/rum-gonic/server"
)

func MiddlewareLogger(log logr.Logger) server.HandlerFunc {
	return func(c *gin.Context) {
		_ = WithLogger(c, log)
	}
}
