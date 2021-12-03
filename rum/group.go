package rum

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tangx/ginbinder"
)

// 结构检查
var _ Handler = (*RouterGroup)(nil)
var _ GroupHandler = (*RouterGroup)(nil)

type RouterGroup struct {
	*gin.RouterGroup

	path     string
	children map[*RouterGroup]bool
	handlers []LogicHandler
}

func NewRouterGroup(path string) *RouterGroup {
	return &RouterGroup{
		path:     path,
		children: make(map[*RouterGroup]bool),
		handlers: make([]LogicHandler, 0),
	}
}

// Output 实现 Handler interface
func (g *RouterGroup) Output(c *gin.Context) (interface{}, error) {
	return nil, nil
}

// GroupHandler 实现 GroupHandler interface
func (g *RouterGroup) getRouterGroup() *RouterGroup {
	return g
}

// Register 添加子 router group 或 logic router
func (g *RouterGroup) Register(handlers ...Handler) {

	if g.handlers == nil {
		g.handlers = make([]LogicHandler, 0)
	}

	for _, handler := range handlers {
		if ghandler, ok := handler.(GroupHandler); ok {
			g.children[ghandler.getRouterGroup()] = true
			continue
		}

		if lhandler, ok := handler.(LogicHandler); ok {
			g.handlers = append(g.handlers, lhandler)
		}
	}

}

// register 遍历子节点并初始化
func (g *RouterGroup) register(parent *gin.RouterGroup) {

	g.RouterGroup = parent.Group(g.path)
	for _, handler := range g.handlers {
		// 通过反射获取 path
		path := handlerPath(handler)

		// 通过断言接口获取 path
		if path == "" {
			h, ok := handler.(PathHandler)
			if !ok {
				continue
			}

			path = h.Path()
		}
		g.RouterGroup.Handle(handler.Method(), path, g.handle(handler))
	}

	for child := range g.children {
		child.register(g.RouterGroup)
	}
}

// handle 处理业务逻辑， 在 gin 中注册路由
func (g *RouterGroup) handle(handler LogicHandler) func(*gin.Context) {

	return func(c *gin.Context) {

		err := ginbinder.ShouldBindRequest(c, handler)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		ret, err := handler.Output(c)
		if err != nil {
			c.JSON(500, err.Error())
			return
		}

		c.JSON(200, ret)
	}
}
