package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.Use(gin.BasicAuth(authUers))

	r.GET("/index", func(c *gin.Context) {
		c.String(200, "index")
	})
	_ = r.Run()
}

var authUers = map[string]string{
	"user1": "tangxin",
}
