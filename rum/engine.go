package rum

import (
	"github.com/gin-gonic/gin"
)

type Engine struct {
	*gin.Engine

	*RouterGroup
}

// Default 默认 gin.Engine
func Default() *Engine {
	r := gin.Default()

	return New(r)
}

// New 使用自定义 gin.Engine
func New(r *gin.Engine) *Engine {

	root := NewRouterGroup("/")

	e := &Engine{
		Engine:      r,
		RouterGroup: root,
	}

	return e
}

// Run 监听 tcp 端口
func (e *Engine) Run(addrs ...string) error {

	// e.RouterGroup.initial(e.RouterGroup)
	e.register()

	return e.Engine.Run(addrs...)
}

func (e *Engine) register() {
	e.RouterGroup.register(&e.Engine.RouterGroup)
}
