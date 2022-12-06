package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-jarvis/rum-gonic/pkg/httpx"
	"github.com/go-jarvis/rum-gonic/server"
)

func main() {
	e := server.Default()
	e.Use(MiddlewarePath)

	// 1. 添加一个服务
	e.Handle(&Index{})

	// 2. 定义 subPath
	sub := server.NewRumPath("/sub")
	sub.Use(MiddlewarePath)
	sub.Handle(&Index{})
	e.AddPath(sub)

	if err := e.Run(":8081"); err != nil {
		panic(err)
	}
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

func MiddlewarePath(c *gin.Context) {
	fmt.Println(c.Request.URL.Path)
}
