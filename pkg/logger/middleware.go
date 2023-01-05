package logger

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-jarvis/logr"
	"github.com/go-jarvis/rum-gonic/server"
	"go.opentelemetry.io/otel/trace"
)

func MiddlewareLogger(log logr.Logger) server.HandlerFunc {
	return func(c *gin.Context) {
		_ = WithLogger(c, log)
	}
}

func MiddlewareLoggerWithSpan() server.HandlerFunc {
	return func(c *gin.Context) {
		log := FromContext(c)

		span := trace.SpanFromContext(c.Request.Context())
		if span.SpanContext().IsValid() {
			log = log.With("trace",
				fmt.Sprintf("%s-%s",
					span.SpanContext().TraceID().String(),
					span.SpanContext().SpanID().String()),
			)
		}

		WithLogger(c, log)
	}
}
