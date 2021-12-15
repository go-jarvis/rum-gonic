package main

import (
	"context"
	"log"
	"time"

	"github.com/go-jarvis/rum-gonic/internal/example/injector/redis"
	"github.com/go-jarvis/rum-gonic/pkg/middlewares"

	"github.com/go-jarvis/rum-gonic/rum"
)

func main() {

	r := rum.Default()
	r.WithContextInjectors(
		redis.WithRedisInject(),
	)

	// 注册中间件
	r.Register(middlewares.DefaultNoCacheIndex())
	// r.Register(auth.AdminUsers)

	// 添加路由组 / 路由
	// r.Register(apis.RouterRoot)
	// r.Register(&index.Index{})

	// r.Register(apis.RouterV0)

	r.StaticFile("/user", "user.html")

	// r.Static("/userindex", "userindex")
	r.Static("/userindex", "/data/gopath/src/github.com/go-jarvis/rum-gonic/userindex")

	/* 启动方式 */
	// 1. 普通方式启动
	normalRun(r)

	// 2. 启动并控制退出
	listenAndServe(r)
}

func normalRun(r *rum.Engine) {
	if err := r.Run(); err != nil {
		panic(err)
	}
}

func listenAndServe(r *rum.Engine) {

	go func() {
		if err := r.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	ctx := context.Background()
	ctx, stop := context.WithTimeout(ctx, 10*time.Second)
	defer stop()

	time.Sleep(30 * time.Second)

	if err := r.Shutdown(ctx); err != nil {
		log.Println("强制关闭 engine")
	}

	log.Println("rum 已经退出")
	// 一分钟后关闭
	time.Sleep(1 * time.Minute)
}
