package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tangx/ginbinder"
)

type HandlerFunc = gin.HandlerFunc

type rumServer struct {
	engine *gin.Engine
	group  *rumRouter
}

func Default() *rumServer {
	e := gin.Default()
	g := e.Group("/")

	group := NewRouter("/")
	group.withGinRG(g)

	return &rumServer{
		engine: e,
		group:  group,
	}
}

func (e *rumServer) Run(addr string) error {
	e.initial()
	return e.engine.Run(addr)
}

func (e *rumServer) initial() {
	e.group.initial()
}

func (e *rumServer) Use(handlers ...HandlerFunc) {
	e.group.Use(handlers...)
}

func (e *rumServer) Handle(handlers ...Operator) {
	e.group.Handle(handlers...)
}

func (e *rumServer) AddRouter(routers ...*rumRouter) {
	e.group.AddRouter(routers...)
}

type rumRouter struct {
	path  string
	ginRG *gin.RouterGroup

	subRouters []*rumRouter

	operators   []Operator
	middlewares []HandlerFunc
}

func NewRouter(path string) *rumRouter {
	return &rumRouter{
		path:       path,
		subRouters: make([]*rumRouter, 0),
	}
}

// withGinRG 添加 gin.RouterGroup
func (rr *rumRouter) withGinRG(rg *gin.RouterGroup) {
	rr.ginRG = rg
}

// initial 初始化自身以及子路由
func (rr *rumRouter) initial() {
	rr.use()
	rr.handle()

	for _, sub := range rr.subRouters {
		subrg := rr.ginRG.Group(sub.path)
		sub.withGinRG(subrg)

		sub.initial()
	}
}

// Use 注册中间件
func (rr *rumRouter) Use(middlewares ...HandlerFunc) {
	// rp.ginRG.Use(handlers...)
	rr.middlewares = append(rr.middlewares, middlewares...)
}

// use 在 initial 调用时， 注册中间件
func (rr *rumRouter) use() {
	rr.ginRG.Use(rr.middlewares...)
}

// Handle 添加业务逻辑
func (rr *rumRouter) Handle(operators ...Operator) {
	rr.operators = append(rr.operators, operators...)
}

// handle 在 initial 调用时， 绑定服务到 gin.RouterGroup
func (rr *rumRouter) handle() {
	for _, oper := range rr.operators {
		op, ok := oper.(APIOperator)
		if !ok {
			continue
		}
		rr.ginRG.Handle(op.Method(), op.Path(), handle(op))
	}
}

// AddRouter 添加子路由
func (rr *rumRouter) AddRouter(groups ...*rumRouter) {
	rr.subRouters = append(rr.subRouters, groups...)
}

// handle 处理业务逻辑
func handle(op Operator) HandlerFunc {
	return func(c *gin.Context) {

		// 参数绑定
		cc := c.Copy()
		err := ginbinder.ShouldBindRequest(cc, op)
		if err != nil {

			c.AbortWithStatusJSON(http.StatusBadRequest, wrapError(err))
			return
		}

		// 业务逻辑处理
		result, err := op.Output(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, wrapError(err))
			return
		}

		switch v := result.(type) {
		case string:
			c.String(http.StatusOK, v)
			return
		}
		c.JSON(http.StatusOK, result)
	}
}

// wrapError 包裹错误
func wrapError(err error) any {
	return map[string]any{
		"error": err.Error(),
	}
}
