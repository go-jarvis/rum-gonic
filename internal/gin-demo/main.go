package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/root", "/data/gopath/src/github.com/tangx/rum-gonic/userindex")

	r.GET("/index", handler)
	_ = r.Run(":8080")
}

var (
	key   = "123"
	value = "bac"
)

func handler(c *gin.Context) {

	// ctx := context.WithValue(c, key, value)

	c.Set(key, value)
	// handler1(c)
	handler2(c)
}

func handler1(c *gin.Context) {

	val := c.MustGet(key)
	fmt.Println(val)
}

func handler2(ctx context.Context) {

	ctx = context.WithValue(ctx, "alibaba", "alimama")

	val := ctx.Value(key)
	fmt.Println(val)

}
