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

/* 嵌套了 httpx.MethodXXX 和 path tag， 以下不需要 */
// func (*Index) Method() string {
// 	return http.MethodGet
// }

// func (*Index) Path() string {
// 	return "/index"
// }
