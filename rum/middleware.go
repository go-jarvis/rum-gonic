package rum

import (
	"github.com/gin-gonic/gin"
)

type MiddlewareOperator interface {
	MiddlewareFunc() HandlerFunc
	Operator
}

// 接口检查
var _ Operator = (*Middleware)(nil)
var _ MiddlewareOperator = (*Middleware)(nil)

type Middleware struct {
	middwareFunc HandlerFunc
	// Operator
}

type HandlerFunc = gin.HandlerFunc

func NewMiddleware(fn HandlerFunc) *Middleware {
	return &Middleware{
		middwareFunc: fn,
	}
}

func (mid *Middleware) MiddlewareFunc() HandlerFunc {
	// fmt.Println("注册中间件咯")
	return mid.middwareFunc
}

func (mid *Middleware) Output(c *gin.Context) (interface{}, error) {
	return nil, nil
}
