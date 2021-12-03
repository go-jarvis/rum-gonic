package main

import (
	"github.com/tangx/rum-gonic/internal/example/apis"
	"github.com/tangx/rum-gonic/internal/example/apis/index"
	"github.com/tangx/rum-gonic/middlewares"
	"github.com/tangx/rum-gonic/rum"
)

func main() {

	r := rum.Default()

	// 注册中间件
	r.Register(middlewares.NoCacheIndex())

	// 添加路由组 / 路由
	r.Register(apis.RouterRoot)
	r.Register(&index.Index{})

	r.Register(apis.RouterV0)

	if err := r.Run(); err != nil {
		panic(err)
	}
}
