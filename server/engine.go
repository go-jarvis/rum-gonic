package server

import (
	"net/http"
	"path/filepath"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-jarvis/rum-gonic/pkg/openapi31"
	"github.com/go-jarvis/rum-gonic/pkg/operator"
	"github.com/go-jarvis/rum-gonic/pkg/reflectx"
	"github.com/tangx/ginbinder"
)

type HandlerFunc = gin.HandlerFunc

type rumServer struct {
	engine *gin.Engine
	router *rumRouter
}

func (e *rumServer) Run(addr string) error {
	e.initial()

	openapi31.Output()

	return e.engine.Run(addr)
}

func (e *rumServer) initial() {
	e.router.initial()
}

func (e *rumServer) Use(handlers ...HandlerFunc) {
	e.router.Use(handlers...)
}

func (e *rumServer) Handle(handlers ...operator.Operator) {
	e.router.Handle(handlers...)
}

func (e *rumServer) AddRouter(routers ...*rumRouter) {
	e.router.AddRouter(routers...)
}

type rumRouter struct {
	path  string
	ginRG *gin.RouterGroup

	subRouters []*rumRouter

	operators   []operator.Operator
	middlewares []HandlerFunc

	// 当前 router 的完全路径
	absolutelyPath string
}

func NewRouter(path string) *rumRouter {
	return &rumRouter{
		path:       path,
		subRouters: make([]*rumRouter, 0),
	}
}

// setGinRG 添加 gin.RouterGroup
func (rr *rumRouter) setGinRG(rg *gin.RouterGroup) *rumRouter {
	rr.ginRG = rg
	return rr
}

// initial 初始化自身以及子路由
func (rr *rumRouter) initial() {
	rr.use()
	rr.handle()

	// 遍历 sub group
	for _, sub := range rr.subRouters {
		subrg := rr.ginRG.Group(sub.path)
		sub.setGinRG(subrg)

		// 设置 sub 的完全路径
		sub.absolutelyPath = filepath.Join(rr.absolutelyPath, sub.path)

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
func (rr *rumRouter) Handle(operators ...operator.Operator) {
	rr.operators = append(rr.operators, operators...)
}

// handle 在 initial 调用时， 绑定服务到 gin.RouterGroup
func (rr *rumRouter) handle() {
	for _, oper := range rr.operators {
		method, path := methodPath(oper)

		rr.ginRG.Handle(method, path, handle(oper))

		abpath := filepath.Join(rr.absolutelyPath, path)

		// 添加 openapi
		openapi31.AddRouter(abpath, method, oper)
	}
}

func methodPath(oper operator.Operator) (method string, path string) {

	if op, ok := oper.(operator.MethodOperator); ok {
		method = op.Method()
	}
	if op, ok := oper.(operator.PathOperator); ok {
		path = op.Path()
	}

	if path != "" {
		return method, path
	}

	rt := reflect.TypeOf(oper)
	rt = reflectx.Deref(rt)
	for i := 0; i < rt.NumField(); i++ {
		ft := rt.Field(i)
		// 取一个
		val, ok := ft.Tag.Lookup("route")
		if ok {
			path = val
			break
		}
	}

	return method, path
}

// AddRouter 添加子路由
func (rr *rumRouter) AddRouter(groups ...*rumRouter) {
	rr.subRouters = append(rr.subRouters, groups...)
}

// handle 处理业务逻辑
func handle(op operator.Operator) HandlerFunc {
	return func(c *gin.Context) {

		op := operator.DeepCopy(op)

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
