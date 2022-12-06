package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tangx/ginbinder"
)

type Server struct {
	engine *gin.Engine
}

func Default() *Server {
	e := gin.Default()

	return &Server{
		engine: e,
	}
}

func (e *Server) Run(addr string) error {
	return e.engine.Run(addr)
}

type HandlerFunc = gin.HandlerFunc

func (e *Server) Use(fns ...HandlerFunc) {
	for _, fn := range fns {
		e.engine.Use(fn)
	}
}

func (e *Server) Handle(handlers ...HanlderOperator) {
	for _, h := range handlers {
		e.handle(h)
	}
}
func (e *Server) handle(handler HanlderOperator) (interface{}, error) {

	e.engine.Handle(handler.Methods(), handler.Path(), handle(handler))
	return nil, nil
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
