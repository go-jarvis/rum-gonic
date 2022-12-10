package operator

import "github.com/gin-gonic/gin"

type APIOperator interface {
	Method() string
	Path() string
}

type Operator interface {
	Output(*gin.Context) (any, error)
}
