# rum-gonic

`rum-gonic` 一个基于 gin 二开的 web 框架。


核心逻辑

1. 将 RouterGroup 抽象成为独立节点， 使用 `RouterGroup.Register(subgroup)` 的方式组合成路由树。 类似 `cobra.Command` 与 `Comamnd.AddCommand(subcommand)` 。

2. 将分来的 `http method`, `uri path` 和 `logic handler` 耦合在一起，成为一个整体。

## demo

```go
package index

import (
	"github.com/gin-gonic/gin"
	"github.com/tangx/rum-gonic/httpx"
)

type Index struct {
	httpx.MethodGet `path:"/index"`
	Name            string `query:"name"`
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
