package rum

import (
	"github.com/gin-gonic/gin"
)

type Engine struct {
	engine *gin.Engine

	*Router
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
		engine: r,
		Router: root,
	}

	return e
}

// Run 监听 tcp 端口
func (e *Engine) Run(addrs ...string) error {

	e.register()

	return e.engine.Run(addrs...)
}

func (e *Engine) register() {
	e.Router.register(&e.engine.RouterGroup)
}
