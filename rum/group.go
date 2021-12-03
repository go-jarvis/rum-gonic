package rum

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tangx/ginbinder"
)

// 接口检查
var _ Operator = (*RouterGroup)(nil)
var _ GroupOperator = (*RouterGroup)(nil)

type GroupOperator interface {
	RouterGroup() *RouterGroup
}

type RouterGroup struct {
	// 当前路由
	path      string
	ginRG     *gin.RouterGroup
	operators []Operator

	// 子路由
	children map[*RouterGroup]bool

	// 接口实现
	Operator
}

func NewRouterGroup(path string) *RouterGroup {
	return &RouterGroup{
		path:      path,
		children:  make(map[*RouterGroup]bool),
		operators: make([]Operator, 0),
	}
}

// RouterGroup 实现 GroupOperator interface
func (r *RouterGroup) RouterGroup() *RouterGroup {
	return r
}

// Register 添加子 router group 或 logic router
func (r *RouterGroup) Register(ops ...Operator) {

	for _, op := range ops {
		// 加入 子路由
		if groupOp, ok := op.(GroupOperator); ok {
			r.children[groupOp.RouterGroup()] = true
			continue
		}

		// 加入 middleware operator 或 logic operator
		if logicOp, ok := op.(Operator); ok {
			r.operators = append(r.operators, logicOp)
		}
	}

}

// register 遍历子节点并初始化
func (r *RouterGroup) register(parent *gin.RouterGroup) {

	// 注册子路由组
	r.ginRG = parent.Group(r.path)

	for _, op := range r.operators {
		// 添加中间件
		if r.use(op) {
			continue
		}

		// 添加 用户逻辑 路由
		r.hanlde(op)
	}

	for child := range r.children {
		child.register(r.ginRG)
	}
}

// use 添加中间件
func (r *RouterGroup) use(op Operator) bool {
	if mid, ok := op.(MiddlewareOperator); ok {
		r.ginRG.Use(mid.MiddlewareFunc())
		return true
	}

	return false
}

// handle 添加路由
func (r *RouterGroup) hanlde(op Operator) bool {

	// 通过反射获取 path
	path := routePath(op)

	// 通过断言接口获取 path
	if path == "" {
		h := op.(PathOperator)
		path = h.Path()
	}

	mop := op.(MethodOperator)

	r.ginRG.Handle(mop.Method(), path, r.handlerfunc(op))
	return true
}

// handlerfunc 处理业务逻辑， 在 gin 中注册路由
func (r *RouterGroup) handlerfunc(op Operator) HandlerFunc {

	return func(c *gin.Context) {

		err := ginbinder.ShouldBindRequest(c, op)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		ret, err := op.Output(c)
		if err != nil {
			c.JSON(500, err.Error())
			return
		}

		c.JSON(200, ret)
	}
}
