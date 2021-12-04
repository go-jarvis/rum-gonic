package main

import (
	"github.com/tangx/rum-gonic/internal/example/apis"
	"github.com/tangx/rum-gonic/internal/example/apis/index"
	"github.com/tangx/rum-gonic/internal/example/injector/redis"
	"github.com/tangx/rum-gonic/middlewares"
	"github.com/tangx/rum-gonic/rum"
)

func main() {

	r := rum.Default()
	r.WithContextInjectors(
		redis.RedisOnlineInjector,
	)

	// 注册中间件
	r.Register(middlewares.NoCacheIndex())

	// 添加路由组 / 路由
	r.Register(apis.RouterRoot)
	r.Register(&index.Index{})

	r.Register(apis.RouterV0)

	r.StaticFile("/user", "user.html")

	// r.Static("/userindex", "userindex")
	r.Static("/userindex", "/data/gopath/src/github.com/tangx/rum-gonic/userindex")

	if err := r.Run(); err != nil {
		panic(err)
	}
}
