package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-jarvis/rum-gonic/pkg/httpx"
	"github.com/go-jarvis/rum-gonic/pkg/logr"
)

type Index struct {
	// httpx.MethodAny `path:"/index.any"`
	httpx.MethodMulti `path:"/index" methods:"GET,HEAD"`
	// httpx.MethodMulti `path:"" methods:"GET,HEAD"`
	Name string `query:"name"`
}

func (index *Index) Output(c *gin.Context) (interface{}, error) {

	return logic(c, index), nil
}

func logic(ctx context.Context, index *Index) map[string]string {

	log := logr.FromContext(ctx)
	log.Info("index logic", "k", "v")

	return map[string]string{
		"code":    "200",
		"message": "index.html",
		"name":    index.Name,
	}
}
