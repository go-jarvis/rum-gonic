package operator

import "github.com/gin-gonic/gin"

type APIOperator interface {
	MethodOperator
	PathOperator
}

type Operator interface {
	Output(*gin.Context) (any, error)
}

type MethodOperator interface {
	Method() string
}

type PathOperator interface {
	Path() string
}
