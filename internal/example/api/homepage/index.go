package homepage

import (
	"github.com/gin-gonic/gin"
	"github.com/go-jarvis/rum-gonic/pkg/httpx"
	"github.com/go-jarvis/rum-gonic/pkg/logger"
	"github.com/go-jarvis/rum-gonic/pkg/operator"
	"github.com/go-jarvis/rum-gonic/server"
)

var IndexRouter = server.NewRouter("")

func init() {
	IndexRouter.Handle(&Index{})
}

var _ operator.APIOperator = &Index{}

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
