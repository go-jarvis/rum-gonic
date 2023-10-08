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
package homepage

import (
	"github.com/gin-gonic/gin"
	"github.com/go-jarvis/rum-gonic/pkg/httpx"
	"github.com/go-jarvis/rum-gonic/pkg/logger"
	"github.com/go-jarvis/rum-gonic/server"
)

var IndexRouter = server.NewRouter("")

func init() {
	IndexRouter.Handle(&Index{})
}

type Index struct {
	httpx.MethodGet `path:"/index/:id"`
	ID              string
}

// Path() 如存在， 优先使用此处的 path 值。
// func (index *Index) Path() string {
// 	return "/index"
// }

func (index *Index) Output(c *gin.Context) (any, error) {

	log := logger.FromContext(c)

	log = log.With("kk", "vv").Start()
	defer log.Stop()

	log.Debug("number %d", 100)
	log.Info("name %s", "index")

	result := map[string]string{
		"name": "zhangsan",
	}
	return result, nil
}
```

## OpenAPI v3.1 Support

**配置 OpenAPI 信息**

```go
	// 初始化内置 openapi reflector
	_ = openapi31.Init()
	// 设置信息
	openapi31.WithOptions(
		openapi31.WithTitle("my-app"),
		openapi31.WithDescription("this's a demo app"),
		openapi31.WithVersion("v0.1.0"),
		openapi31.WithFile("openapi2.yaml"),
	)

	// 输出 openapi31 信息
	// 在 rum server 中会被默认调用
	openapi31.Output()
```

**路由添加 OpenAPI 信息**

```go
type Index struct {
	httpx.MethodGet
	ID   string   `path:"id" example:"xxx-xxxx"`
	Name []string `query:"name" example:"Mike Jackson"`
}

type User struct {
	httpx.MethodPost
	Class string `path:"class" example:"yyy"`

	User struct {
		Name string `json:"name" example:"joe bidden"`
		Age  int    `json:"age" example:"100"`
	} `body:"body" mime:"json"`
}
```

1. `path, query, cookie, header, json` 等表示 **参数位置与名字**， 参考 [ginbinder 参数规则](https://github.com/tangx/ginbinder)， 同时也是 [openapi 规则](https://github.com/swaggest/openapi-go#features)
2. `example` 表示 **默认值**