package main

import (
	"github.com/tangx/rum-gonic/internal/example/apis"
	"github.com/tangx/rum-gonic/internal/example/apis/index"
	"github.com/tangx/rum-gonic/rum"
)

func main() {

	r := rum.Default()

	r.Register(apis.RouterRoot)
	r.Register(&index.Index{})

	r.Register(apis.RouterV0)

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
