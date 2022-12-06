package server

import "github.com/gin-gonic/gin"

type APIOperator interface {
	Methods() string
	Path() string
	Operator
}

type Operator interface {
	Output(*gin.Context) (any, error)
}
