package failed

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-jarvis/rum-gonic/pkg/httpx"
)

type CreateUser struct {
	httpx.MethodPost `path:"/create-user"`

	Data  CreateUserParams `body:"body"`
	Sleep int              `query:"sleep"`
}

type CreateUserParams struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (u *CreateUser) Output(c *gin.Context) (interface{}, error) {

	fmt.Println("before:", u)

	d := time.Duration(u.Sleep) * time.Second
	time.Sleep(d)

	fmt.Println("after:", u)

	return u, nil
}
