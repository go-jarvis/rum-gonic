package main

import (
	"github.com/go-jarvis/rum-gonic/pkg/middlewares/trace"
	"github.com/go-jarvis/rum-gonic/rum"
)

func main() {
	r := rum.Default()

	r.Use(
		trace.TraceSpanExtractMiddleware,
		trace.TraceSpanInjectMiddleware,
		trace.WithSpanLoggerMiddleware,
	)

	r.Register(&Index{})

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
