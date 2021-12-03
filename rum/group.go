package rum

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tangx/ginbinder"
)

// 接口检查
// var _ Operator = (*Router)(nil)
var _ GroupOperator = (*Router)(nil)

type Router struct {
	// 当前路由
	path      string
	ginRG     *gin.RouterGroup
	operators []Operator

	// 子路由
	children map[*Router]bool

	// 接口实现
	Operator
}

func NewRouterGroup(path string) *Router {
	return &Router{
		path:      path,
		children:  make(map[*Router]bool),
		operators: make([]Operator, 0),
	}
}

// getRouter 实现 GroupOperator interface
func (r *Router) getRouter() *Router {
	return r
}

// Register 添加子 router group 或 logic router
func (r *Router) Register(ops ...Operator) {

	// if r.operators == nil {
	// 	r.operators = make([]LogicOperator, 0)
	// }

	for _, op := range ops {
		if groupOp, ok := op.(GroupOperator); ok {
			r.children[groupOp.getRouter()] = true
			continue
		}

		if logicOp, ok := op.(LogicOperator); ok {
			r.operators = append(r.operators, logicOp)
		}
	}

}

// register 遍历子节点并初始化
func (r *Router) register(parent *gin.RouterGroup) {

	r.ginRG = parent.Group(r.path)
	for _, op := range r.operators {
		// 通过反射获取 path
		path := routePath(op)
		// 通过断言接口获取 path
		if path == "" {
			h, ok := op.(PathOperator)
			if !ok {
				continue
			}

			path = h.Path()
		}

		mop, ok := op.(MethodOperator)
		if !ok {
			continue
		}
		r.ginRG.Handle(mop.Method(), path, r.handle(op))
	}

	for child := range r.children {
		child.register(r.ginRG)
	}
}

// handle 处理业务逻辑， 在 gin 中注册路由
func (r *Router) handle(op Operator) func(*gin.Context) {

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
