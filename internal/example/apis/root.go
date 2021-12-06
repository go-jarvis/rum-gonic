package apis

import (
	"github.com/tangx/rum-gonic/internal/example/apis/index"
	"github.com/tangx/rum-gonic/internal/example/apis/pingpong"
	"github.com/tangx/rum-gonic/internal/example/auth"
	"github.com/tangx/rum-gonic/rum"
)

var (
	RouterRoot = rum.NewRouterGroup("/rum")
	RouterV0   = rum.NewRouterGroup("/v0")
)

func init() {

	RouterRoot.Register(RouterV0)

	{
		RouterV0.Register(auth.AdminUsersMiddelware)
		// rum handler mode
		RouterV0.Register(&pingpong.PingPong{})
		RouterV0.Register(&index.Index{})
		RouterV0.Static("/user", "dist")
	}
}
