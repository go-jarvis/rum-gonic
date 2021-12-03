package pingpong

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingPong struct {
	Name string `uri:"name"`
	Age  int    `query:"age"`
}

func (pp *PingPong) Path() string {
	return "/ping/:name"
}

func (*PingPong) Method() string {
	return http.MethodGet
}

func (pp *PingPong) Output(c *gin.Context) (interface{}, error) {
	return *pp, nil
}
