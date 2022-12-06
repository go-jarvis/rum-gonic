package api

import (
	"github.com/go-jarvis/rum-gonic/internal/example/api/homepage"
	"github.com/go-jarvis/rum-gonic/internal/example/api/sub"
	"github.com/go-jarvis/rum-gonic/server"
)

var RootApp = server.NewRumPath("/app")

func init() {
	RootApp.AddPath(homepage.IndexRouter)
	RootApp.AddPath(sub.SubRouter)
}
