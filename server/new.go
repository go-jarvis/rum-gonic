package server

import (
	"github.com/gin-gonic/gin"
)

func Default() *RumServer {
	e := New()

	e.Use(gin.Logger(), gin.Recovery())

	return e
}

func New() *RumServer {
	e := gin.New()

	rg := e.Group("/")
	router := NewRouter("/").setGinRG(rg)

	return &RumServer{
		engine: e,
		router: router,
	}
}
