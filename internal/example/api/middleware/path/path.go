package path

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func MiddlewarePath(c *gin.Context) {
	fmt.Println(c.Request.URL.Path)
}
