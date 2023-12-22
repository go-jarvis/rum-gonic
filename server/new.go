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

	rs := &RumServer{}
	rs.Initialize()

	return rs
}
