package rum

import "github.com/gin-gonic/gin"

type PathHandler interface {
	Path() string
}

type LogicHandler interface {
	Method() string
	Handler
}

type GroupHandler interface {
	getRouterGroup() *RouterGroup
}

type Handler interface {
	Output(*gin.Context) (interface{}, error)
}
