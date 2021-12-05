package rum

import (
	"github.com/gin-gonic/gin"
)

type Engine struct {
	engine *gin.Engine

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
		engine:      r,
		RouterGroup: root,
	}

	return e
}

// Run 监听 tcp 端口
func (e *Engine) Run(addrs ...string) error {

	e.register()

	return e.engine.Run(addrs...)
}

func (e *Engine) register() {
	root := &RouterGroup{
		ginRG: &e.engine.RouterGroup,
	}
	e.RouterGroup.register(root)
}

// WithContextCompose 添加注入控制器
func (e *Engine) WithContextCompose(injectors ...ContextInjectorFunc) *Engine {
	if contextInjectors == nil {
		contextInjectors = make([]ContextInjectorFunc, 0)
	}
	contextInjectors = append(contextInjectors, injectors...)

	return e
}
