package index

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-jarvis/rum-gonic/internal/example/injector/redis"
	"github.com/go-jarvis/rum-gonic/pkg/httpx"
	"github.com/go-jarvis/statuserrors"
)

type Index struct {
	// httpx.MethodAny `path:"/index.any"`
	httpx.MethodMulti `path:"/index" methods:"GET,HEAD"`
	// httpx.MethodMulti `path:"" methods:"GET,HEAD"`
	Name string `query:"name"`
}

func (index *Index) Output(c *gin.Context) (interface{}, error) {
	msg := c.Query("e")
	if msg == "nil" {
		err := statuserrors.New(http.StatusBadRequest, "invalid request")
		return nil, err
	}

	if msg == "data" {
		err := statuserrors.New(http.StatusBadRequest, "invalid request")
		return "this is user define output data, not from error", err
	}

	return logic(c, index), nil
}

func logic(ctx context.Context, index *Index) map[string]string {

	ra := redis.FromRedisAgentOnline(ctx)

	return map[string]string{
		"redis-agent": ra.ServerAddr(),
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
