package index

import (
	"github.com/gin-gonic/gin"
	"github.com/tangx/rum-gonic/httpx"
)

type Index struct {
	httpx.MethodGet `path:"/index"`
}

func (index *Index) Output(c *gin.Context) (interface{}, error) {
	return map[string]string{
		"code":    "200",
		"message": "index.html",
	}, nil
}
