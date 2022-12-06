package api

import (
	"github.com/go-jarvis/rum-gonic/internal/example/api/homepage"
	"github.com/go-jarvis/rum-gonic/internal/example/api/sub"
	"github.com/go-jarvis/rum-gonic/server"
)

var RootApp = server.NewRouter("/app")

func init() {
	RootApp.AddRouter(homepage.IndexRouter)
	RootApp.AddRouter(sub.SubRouter)
}
