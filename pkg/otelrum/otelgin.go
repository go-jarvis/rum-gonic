// OpenTelemetry for Gin
package otelrum

import (
	"github.com/gin-gonic/gin"
	"github.com/go-jarvis/rum-gonic/server"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func Middleware(appname string, opts ...otelgin.Option) server.HandlerFunc {
	return func(ctx *gin.Context) {
		otelgin.Middleware(appname, opts...)
	}
}
