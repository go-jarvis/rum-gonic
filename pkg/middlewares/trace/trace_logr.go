package trace

import (
	"github.com/gin-gonic/gin"
	"github.com/go-jarvis/rum-gonic/pkg/logr"
)

func WithSpanLoggerMiddleware(c *gin.Context) {

	span := TraceSpanFromContext(c)

	if span.SpanContext().IsValid() {
		log := logr.FromContext(c)

		log = log.With(
			"spanId", span.SpanContext().SpanID(),
			"traceId", span.SpanContext().TraceID(),
		)
		// re-inject
		_ = logr.WithContext(c, log)
	}

	c.Next()
}
