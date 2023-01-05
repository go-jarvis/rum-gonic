package server

import (
	"github.com/gin-gonic/gin"
)

func Default() *rumServer {
	e := New()

	e.Use(gin.Logger(), gin.Recovery())

	return e
}

func New() *rumServer {
	e := gin.New()

	rg := e.Group("/")
	router := NewRouter("/").withGinRG(rg)

	return &rumServer{
		engine: e,
		router: router,
	}
}
