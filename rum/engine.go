package rum

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Engine struct {
	engine *gin.Engine

	*RouterGroup

	srv *http.Server
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

// Run 监听 tcp 端口
func (e *Engine) Run(addrs ...string) error {

	e.register()

	return e.engine.Run(addrs...)
}

// ListenAndServe 启动服务， 可以配合 Shutdown(ctx) 自定义退出
func (e *Engine) ListenAndServe(addrs ...string) error {
	e.register()

	if e.srv == nil {
		e.srv = &http.Server{
			Addr:    addr(addrs...),
			Handler: e.engine,
		}
	}

	err := e.srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

// Shutdown 用户自定义退出。 如果不是使用 ListenAndServe 启动则不能使用。
func (e *Engine) Shutdown(ctx context.Context) error {
	if e.srv == nil {
		return errors.New(NotRunWithListenAndServe)
	}

	return e.srv.Shutdown(ctx)
}

func addr(addrs ...string) string {
	if len(addrs) == 0 {
		return ":8080"
	}

	return addrs[0]
}
