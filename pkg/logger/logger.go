package logger

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func NewSLoggerHandler(c *gin.Context) {

	logr := NewJSONLogger()

	logr = logr.With(
		"traceId", c.Value("traceId"),
		"spanId", c.Value("spanId"),
		"pspanId", c.Value("pspanId"),
	)

	c.Set(Key_SLogger, logr)
	c.Next()

}

var (
	Key_SLogger = "slogger.asdfjkaldfjlsdjf.dsfjalsdkfjalsjdf"
)

func FromContext(ctx context.Context) *slog.Logger {
	logr := ctx.Value(Key_SLogger)

	if logr != nil {
		return logr.(*slog.Logger)
	}

	return NewJSONLogger()
}

func NewJSONLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout))
}
