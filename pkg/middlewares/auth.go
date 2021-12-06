package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/tangx/rum-gonic/rum"
)

func BasicAuth(accounts gin.Accounts) rum.MiddlewareOperator {
	return rum.NewMiddleware(
		gin.BasicAuth(accounts),
	)
}

func BasicAuthRealm(accounts gin.Accounts, realm string) rum.MiddlewareOperator {
	return rum.NewMiddleware(
		gin.BasicAuthForRealm(accounts, realm),
	)
}
