package auth

import "github.com/go-jarvis/rum-gonic/pkg/middlewares"

var AdminUsersMiddelware = middlewares.BasicAuth(map[string]string{
	"user1": "tangxin",
})
