package rum

import "github.com/gin-gonic/gin"

type Operator interface {
	Output(*gin.Context) (interface{}, error)
}

type PathOperator interface {
	Path() string
}

type LogicOperator interface {
	Operator
	MethodOperator
}

type GroupOperator interface {
	getRouter() *Router
}

type MethodOperator interface {
	Method() string
}
