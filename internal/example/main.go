package main

import (
	"github.com/go-jarvis/logr"
	"github.com/go-jarvis/rum-gonic/internal/example/api"
	"github.com/go-jarvis/rum-gonic/pkg/logger"
	"github.com/go-jarvis/rum-gonic/server"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

var tracer = otel.Tracer("gin-server")

func initTracer() (*sdktrace.TracerProvider, error) {
	// exporter, err := stdout.New(stdout.WithPrettyPrint())
	// if err != nil {
	// 	return nil, err
	// }
	tp := sdktrace.NewTracerProvider(
	// sdktrace.WithSampler(sdktrace.AlwaysSample()),
	// sdktrace.WithBatcher(exporter),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp, nil
}

func main() {
	_, _ = initTracer()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer func() {
	// 	if err := tp.Shutdown(context.Background()); err != nil {
	// 		log.Printf("Error shutting down tracer provider: %v", err)
	// 	}
	// }()

	log := logr.New(logr.Config{
		Level: "info",
	})

	e := server.Default()
	e.Use(otelgin.Middleware("my-rum-server-example"))
	e.Use(logger.MiddlewareLogger(log))
	e.Use(logger.MiddlewareLoggerWithSpan())

	e.AddRouter(api.RootApp)
	if err := e.Run(":8081"); err != nil {
		panic(err)
	}
}
