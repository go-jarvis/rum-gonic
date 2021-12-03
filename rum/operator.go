package rum

import (
	"github.com/gin-gonic/gin"
)

type Operator interface {
	Output(*gin.Context) (interface{}, error)
}

type PathOperator interface {
	Path() string
}

type MethodOperator interface {
	Method() string
}
