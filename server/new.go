package server

import (
	"github.com/gin-gonic/gin"
	"github.com/go-jarvis/rum-gonic/pkg/opentelemetry"
)

func Default() *rumServer {
	e := New()

	e.Use(gin.Logger(), gin.Recovery())
	e.Use(
		opentelemetry.TraceSpanExtractMiddleware,
		opentelemetry.TraceSpanInjectMiddleware,
	)

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
