package main

import (
	"github.com/go-jarvis/rum-gonic/internal/example/api"
	"github.com/go-jarvis/rum-gonic/internal/example/api/middleware/path"
	"github.com/go-jarvis/rum-gonic/server"
)

func main() {
	e := server.Default()
	e.Use(path.MiddlewarePath)

	e.AddPath(api.RootApp)
	if err := e.Run(":8081"); err != nil {
		panic(err)
	}
}
