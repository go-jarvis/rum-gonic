package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.Static("/root", "/data/gopath/src/github.com/tangx/rum-gonic/userindex")
	_ = r.Run(":8080")
}
