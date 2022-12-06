package homepage

import (
	"github.com/gin-gonic/gin"
	"github.com/go-jarvis/rum-gonic/pkg/httpx"
	"github.com/go-jarvis/rum-gonic/server"
)

var IndexRouter = server.NewRumPath("")

func init() {
	IndexRouter.Handle(&Index{})
}

var _ server.APIOperator = &Index{}

type Index struct {
	httpx.MethodGet
}

func (index *Index) Path() string {
	return "/index"
}
func (index *Index) Methods() string {
	return index.Method()
}

func (index *Index) Output(c *gin.Context) (any, error) {

	result := map[string]string{
		"name": "zhangsan",
	}
	return result, nil
}
