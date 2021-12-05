package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/tangx/rum-gonic/rum"
)

func NoCacheIndex() *rum.Middleware {

	pages := []string{"/", "/index", "index.html"}

	return NoCahe(pages)
}

func NoCahe(pages []string) *rum.Middleware {

	mid := func(c *gin.Context) {
		path := c.Request.URL.Path

		for _, page := range pages {
			if page == path {
				c.Header("Cache-Control", "no-cache")
				break
			}
		}
	}

	return rum.NewMiddleware(mid)
}
