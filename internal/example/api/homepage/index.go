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
	httpx.MethodGet `route:"/index/:id"`
	ID              string
}

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
