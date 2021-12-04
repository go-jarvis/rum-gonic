package index

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tangx/rum-gonic/httpx"
	"github.com/tangx/rum-gonic/internal/example/injector/redis"
)

type Index struct {
	httpx.MethodGet `path:"/index"`
	Name            string `query:"name"`
}

func (index *Index) Output(c *gin.Context) (interface{}, error) {

	return logic(c, index), nil
}

func logic(ctx context.Context, index *Index) map[string]string {

	ra := redis.FromRedisAgentOnline(ctx)

	return map[string]string{
		"redis-agent": fmt.Sprintf("%s:%d", ra.Addr, ra.Port),
		"code":        "200",
		"message":     "index.html",
		"name":        index.Name,
	}
}

/* 嵌套了 httpx.MethodXXX 和 path tag， 以下不需要 */
// func (*Index) Method() string {
// 	return http.MethodGet
// }

// func (*Index) Path() string {
// 	return "/index"
// }
