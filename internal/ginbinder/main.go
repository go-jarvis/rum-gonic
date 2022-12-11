package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-jarvis/rum-gonic/pkg/httpx"
	"github.com/go-jarvis/rum-gonic/server"
)

func main() {

	r := server.Default()
	r.Handle(&Index{})

	err := r.Run(":8089")
	if err != nil {
		panic(err)
	}
}

type Index struct {
	httpx.MethodGet
	Name []string `query:"name"`
}

func (*Index) Method() string {
	return "GET"
}
func (*Index) Path() string {
	return "/"
}

func (index *Index) Output(c *gin.Context) (any, error) {

	return index.Name, nil
}
