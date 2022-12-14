package main

import (
	"github.com/go-jarvis/logr"
	"github.com/go-jarvis/rum-gonic/internal/example/api"
	"github.com/go-jarvis/rum-gonic/internal/example/api/middleware/path"
	"github.com/go-jarvis/rum-gonic/pkg/logger"
	"github.com/go-jarvis/rum-gonic/server"
)

func main() {
	log := logr.New(logr.Config{
		Level: "info",
	})

	e := server.Default()
	e.Use(path.MiddlewarePath)
	e.Use(logger.MiddlewareLogger(log))

	e.AddRouter(api.RootApp)
	if err := e.Run(":8081"); err != nil {
		panic(err)
	}
}
