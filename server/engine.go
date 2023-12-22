package server

import (
	"context"
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

// RumServer 定义 rum server
type RumServer struct {
	Listen string `env:""`

	engine *gin.Engine
	router *rumRouter
	ctx    context.Context
}

func (rs *RumServer) SetDefaults() {
	if rs.ctx == nil {
		rs.ctx = context.Background()
	}

	if rs.Listen == "" {
		rs.Listen = ":8080"
	}
}

func (rs *RumServer) Initialize() {
	e := gin.New()

	rg := e.Group("/")
	router := NewRouter("/").setGinRG(rg)

	rs.engine = e
	rs.router = router
}

// WithContext 设置 context
func (rs *RumServer) WithContext(ctx context.Context) {
	rs.ctx = ctx
}

// Context 获取 context
func (rs *RumServer) Context() context.Context {
	return rs.ctx
}

// injectContext 将 context 注入到 gin.Context.Request 中
func (rs *RumServer) injectContext() {
	h := func(c *gin.Context) {
		r2 := c.Request.WithContext(rs.ctx)
		c.Request = r2
	}

	rs.Use(h)
}

// Run 启动服务
func (rs *RumServer) Run(addr ...string) error {

	rs.injectContext()
	rs.initial()

	openapi31.Output()

	if len(addr) != 0 {
		rs.Listen = addr[0]
	}

	return rs.engine.Run(rs.Listen)
}

// initial 初始化
func (rs *RumServer) initial() {
	rs.router.initial()
}

// Use 注册中间件
func (rs *RumServer) Use(handlers ...HandlerFunc) {
	rs.router.Use(handlers...)
}

// Handle 添加业务逻辑
func (rs *RumServer) Handle(handlers ...operator.Operator) {
	rs.router.Handle(handlers...)
}

// AddRouter 添加子路由
func (rs *RumServer) AddRouter(routers ...*rumRouter) {
	rs.router.AddRouter(routers...)
}

// rumRouter 路由组
type rumRouter struct {
	path  string
	ginRG *gin.RouterGroup

	subRouters []*rumRouter

	operators   []operator.Operator
	middlewares []HandlerFunc

	// 当前 router 的完全路径
	absolutelyPath string
}

// NewRouter 创建路由组
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
