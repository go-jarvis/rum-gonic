package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-jarvis/rum-gonic/rum"
)

// DefaultCorsPolicy 默认跨域规则， 所有来源
func DefaultCorsPolicy() rum.MiddlewareOperator {
	return CorsPolicy("*")
}

// CorsPolicy 允许跨域来源
// example, sorigin = https://developer.mozilla.org
func CorsPolicy(origin string) rum.MiddlewareOperator {

	cors := func(c *gin.Context) {
		method := c.Request.Method
		if method != "" {
			c.Header("Access-Control-Allow-Origin", origin) // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization,X-Token")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}

	return rum.NewMiddleware(cors)
}
