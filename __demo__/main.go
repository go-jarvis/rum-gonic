package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-jarvis/rum-gonic/pkg/httpx"
	"github.com/go-jarvis/rum-gonic/server"
)

func main() {
	r := server.Default()

	r.Handle(&Index{})

	if err := r.Run(":8081"); err != nil {
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
	// err := errors.New("name error")
	return result, nil
}
