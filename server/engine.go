package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tangx/ginbinder"
)

type HandlerFunc = gin.HandlerFunc

type rumServer struct {
	engine *gin.Engine
	group  *rumRouterGroup
}

func Default() *rumServer {
	e := gin.Default()
	rg := e.Group("/")

	return &rumServer{
		engine: e,
		group: &rumRouterGroup{
			path:  "/",
			group: rg,
		},
	}
}

func (e *rumServer) Run(addr string) error {
	return e.engine.Run(addr)
}

func (e *rumServer) Use(handlers ...HandlerFunc) {
	e.group.Use(handlers...)
}

func (e *rumServer) Handle(handlers ...Operator) {
	e.group.Handle(handlers...)
}

type rumRouterGroup struct {
	path  string
	group *gin.RouterGroup
}

// func newRumRouterGroup(path string) *rumRouterGroup {
// 	return &rumRouterGroup{
// 		path: path,
// 	}
// }

func (rg *rumRouterGroup) Use(handlers ...HandlerFunc) {
	rg.group.Use(handlers...)
}

func (rg *rumRouterGroup) Handle(operators ...Operator) {
	for _, oper := range operators {
		op, ok := oper.(APIOperator)
		if !ok {
			continue
		}

		rg.group.Handle(op.Methods(), op.Path(), handle(op))
	}
}

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

func wrapError(err error) any {
	return map[string]any{
		"error": err.Error(),
	}
}
