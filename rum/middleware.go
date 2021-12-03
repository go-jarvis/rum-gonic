package rum

import "github.com/gin-gonic/gin"

type MiddlewareOperator interface {
	Operator
}

func NewMiddleware(fn gin.HandlerFunc) MiddlewareOperator {

	return nil
}

type Middleware struct {
	Operator
}
