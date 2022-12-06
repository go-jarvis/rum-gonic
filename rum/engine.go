package rum

import "github.com/gin-gonic/gin"

type Engine struct {
	engine *gin.Engine
}

func Default() *Engine {
	e := gin.Default()

	return &Engine{
		engine: e,
	}
}

func (e *Engine) Run(addr string) error {
	return e.engine.Run(addr)
}
