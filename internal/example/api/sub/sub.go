package sub

import (
	"github.com/go-jarvis/rum-gonic/internal/example/api/homepage"
	"github.com/go-jarvis/rum-gonic/server"
)

var SubRouter = server.NewRumPath("/sub")

func init() {
	SubRouter.AddPath(homepage.IndexRouter)
}
