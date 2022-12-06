package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tangx/ginbinder"
)

type HandlerFunc = gin.HandlerFunc

type rumServer struct {
	engine *gin.Engine
	group  *rumPath
}

func Default() *rumServer {
	e := gin.Default()
	g := e.Group("/")

	group := NewRumPath("/")
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

func (e *rumServer) AddPath(paths ...*rumPath) {
	e.group.AddPath(paths...)
}

type rumPath struct {
	path  string
	ginRG *gin.RouterGroup

	subPaths []*rumPath

	operators   []Operator
	middlewares []HandlerFunc
}

func NewRumPath(path string) *rumPath {
	return &rumPath{
		path:     path,
		subPaths: make([]*rumPath, 0),
	}
}

// withGinRG 添加 gin.RouterGroup
func (rp *rumPath) withGinRG(rg *gin.RouterGroup) {
	rp.ginRG = rg
}

// initial 初始化自身以及子路由
func (rp *rumPath) initial() {
	rp.use()
	rp.handle()

	for _, sub := range rp.subPaths {
		subrg := rp.ginRG.Group(sub.path)
		sub.withGinRG(subrg)

		sub.initial()
	}
}

// Use 注册中间件
func (rp *rumPath) Use(middlewares ...HandlerFunc) {
	// rp.ginRG.Use(handlers...)
	rp.middlewares = append(rp.middlewares, middlewares...)
}

// use 在 initial 调用时， 注册中间件
func (rp *rumPath) use() {
	rp.ginRG.Use(rp.middlewares...)
}

// Handle 添加业务逻辑
func (rp *rumPath) Handle(operators ...Operator) {
	rp.operators = append(rp.operators, operators...)
}

// handle 在 initial 调用时， 绑定服务到 gin.RouterGroup
func (rp *rumPath) handle() {
	for _, oper := range rp.operators {
		op, ok := oper.(APIOperator)
		if !ok {
			continue
		}
		rp.ginRG.Handle(op.Methods(), op.Path(), handle(op))
	}
}

// AddPath 添加子路由
func (rg *rumPath) AddPath(groups ...*rumPath) {
	rg.subPaths = append(rg.subPaths, groups...)
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
