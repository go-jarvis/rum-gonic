package pingpong

import (
	"github.com/gin-gonic/gin"
	"github.com/tangx/rum-gonic/httpx"
)

type PingPong struct {
	httpx.MethodGet
	Name string `uri:"name"`
	Age  int    `query:"age"`
}

func (pp *PingPong) Path() string {
	return "/ping/:name"
}
func (pp *PingPong) Output(c *gin.Context) (interface{}, error) {
	// return "pong", nil
	return *pp, nil
}
