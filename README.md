# rum-gonic

`rum-gonic` 一个基于 gin 二开的 web 框架。


## 核心思路

1. 将 RouterGroup 抽象成为独立节点， 使用 `RouterGroup.Register(subgroup)` 的方式组合成路由树。 类似 `cobra.Command` 与 `Comamnd.AddCommand(subcommand)` 。

2. 将分来的 `http method`, `uri path` 和 `logic handler` 耦合在一起，成为一个整体。

3. 整合了 `github.com/tangx/ginbinder` 对 logic struct 进行变量赋值。

## 实现逻辑

在 `/rum/operator.go` 定义了数种 Operator interface。

在 `/rum/group.go` 中的 `RouterGroup.Register(...)` 和 `RouterGroup.register(...)` 进行多次 **接口断言** 进行不同 operator 的逻辑处理。 因此可以使用 `RouterGroup.Register()` 这 **一个方法** 注册多种 Operator `(Group, Handler, Middleware)` 。


## demo

完整 example [main.go](/internal/example/main.go)


将 method, path, handler 组合在一起代码如下

```go
package index

import (
	"github.com/gin-gonic/gin"
	"github.com/tangx/rum-gonic/httpx"
)

type Index struct {
	// httpx.MethodAny `path:"/index.any"`
	httpx.MethodMulti `path:"/index" methods:"GET,HEAD"`
	Name              string `query:"name"`
}

func (index *Index) Output(c *gin.Context) (interface{}, error) {
	return logic(index), nil
}

func logic(index *Index) map[string]string {
	return map[string]string{
		"code":    "200",
		"message": "index.html",
		"name":    index.Name,
	}
}
```
