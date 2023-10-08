package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-jarvis/rum-gonic/pkg/httpx"
	"github.com/go-jarvis/rum-gonic/pkg/openapi31"

	"github.com/go-jarvis/rum-gonic/server"
)

func main() {

	// 初始化内置 openapi reflector
	_ = openapi31.Init()
	// 设置信息
	openapi31.WithOptions(
		openapi31.WithTitle("my-app"),
		openapi31.WithDescription("this's a demo app"),
		openapi31.WithVersion("v0.1.0"),
		openapi31.WithFile("openapi2.yaml"),
	)

	r := server.Default()
	r.Handle(&Index{})

	ng := server.NewRouter("/user-group")
	ng.Handle(&User{})

	r.AddRouter(ng)

	err := r.Run(":8089")
	if err != nil {
		panic(err)
	}
}

type Index struct {
	httpx.MethodGet
	ID   string   `path:"id" example:"xxx-xxxx"`
	Name []string `query:"name" example:"Mike Jackson"`
}

func (*Index) Method() string {
	return http.MethodGet
}
func (*Index) Path() string {
	return "/:id"
}

func (index *Index) Output(c *gin.Context) (any, error) {
	return index.Name, nil
}

type User struct {
	httpx.MethodPost
	Class string `path:"class" example:"yyy"`

	User struct {
		Name string `json:"name" example:"joe bidden"`
		Age  int    `json:"age" example:"100"`
	} `body:"body" mime:"json"`
}

func (*User) Path() string {
	return "/:class"
}

func (user *User) Output(c *gin.Context) (any, error) {
	return user.Class, nil
}
