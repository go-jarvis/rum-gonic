package rum

import "github.com/gin-gonic/gin"

type PathOperator interface {
	Path() string
}

type LogicOperator interface {
	Method() string
	Operator
}

type GroupOperator interface {
	getRouterGroup() *Router
}

type Operator interface {
	Output(*gin.Context) (interface{}, error)
}
