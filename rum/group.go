package rum

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-jarvis/rum-gonic/pkg/httpx"
	"github.com/go-jarvis/statuserrors"
	"github.com/tangx/ginbinder"
)

// 接口检查
var _ Operator = (*RouterGroup)(nil)
var _ GroupOperator = (*RouterGroup)(nil)

type GroupOperator interface {
	routerGroup() *RouterGroup
}

type RouterGroup struct {
	// 当前路由
	path      string
	ginRG     *gin.RouterGroup
	operators []Operator

	// 子路由
	children map[*RouterGroup]bool

	// 其他
	parent *RouterGroup

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
func (r *RouterGroup) routerGroup() *RouterGroup {
	return r
}

// Register 添加子 router group ,  logic , middleware
func (r *RouterGroup) Register(ops ...Operator) {

	for _, op := range ops {
		// 加入 子路由
		if groupOp, ok := op.(GroupOperator); ok {
			r.children[groupOp.routerGroup()] = true
			continue
		}

		// 加入 middleware operator 或 logic operator
		if logicOp, ok := op.(Operator); ok {
			r.operators = append(r.operators, logicOp)
		}
	}
}

// register 遍历子节点并初始化
func (r *RouterGroup) register(parent *RouterGroup) {
	if r == parent {
		panic("自己不能注册自己")
	}
	r.parent = parent

	// 注册子路由组
	r.ginRG = r.parent.ginRG.Group(r.path)

	for _, op := range r.operators {
		// 添加中间件
		if r.use(op) {
			continue
		}

		// 添加 用户逻辑 路由
		r.hanlde(op)
	}

	for child := range r.children {
		child.register(r)
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

	for _, method := range httpx.UnmarshalMethods(mop.Method()) {
		method = strings.TrimSpace(method)
		// 有错就要报，免得找不到
		// if len(method) == 0 {
		// 	continue
		// }
		r.ginRG.Handle(method, path, r.handlerfunc(op))
	}

	return true
}

// handlerfunc 处理业务逻辑， 在 gin 中注册路由
func (r *RouterGroup) handlerfunc(op Operator) HandlerFunc {

	return func(c *gin.Context) {

		err := ginbinder.ShouldBindRequest(c, op)
		if err != nil {
			err = statuserrors.Wrap(err, http.StatusBadRequest, BindingRequestError)
			r.output(c, nil, err)
			return
		}

		// inject 注入
		for _, injector := range contextInjectors {
			_ = injector(c)
		}

		ret, err := op.Output(c)

		// 检测是否在 operator 已经中止， 例如 StaticFile 服务
		if c.IsAborted() {
			return
		}

		r.output(c, ret, err)
	}
}

func (r *RouterGroup) addOperators(ops ...Operator) {
	r.operators = append(r.operators, ops...)
}

// output give response code and data
// content-type is text/plain if data is string type, or content-type
// is application/json by default. maybe it will support more
// content types in feture.
func (r *RouterGroup) output(c *gin.Context, data interface{}, err error) {
	code, data := extract(data, err)

	switch ret := data.(type) {
	case string:
		c.String(code, ret)
	default:
		c.JSON(code, ret)
	}

}

// extract return http status code and output message,
// no matter if error is nil,
// if data is not nil, using data for output message, otherwise trying to use
// error Error() message as output message
func extract(data interface{}, err error) (code int, result interface{}) {

	code, result = extractError(err)

	if data != nil {
		return code, data
	}

	return code, result
}

// extractError return http status code and error message
// if err is not a status error, try to return statuserrors status code
// and error message.
func extractError(err error) (code int, msg string) {
	if err == nil {
		return http.StatusOK, ""
	}

	if e, ok := err.(statuserrors.StatusError); ok {
		return e.StatusCode(), err.Error()
	}

	sterr := statuserrors.New(statuserrors.StatusUnknownError, err.Error())
	return statuserrors.StatusUnknownError, sterr.Error()
}
