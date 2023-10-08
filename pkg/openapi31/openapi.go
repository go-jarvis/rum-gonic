package openapi31

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/swaggest/openapi-go/openapi31"
)

var reflector *openapi31.Reflector

// New 初始化一个 OpenAPI Reflector
func New(title, version, desc string) {
	// 初始化
	reflector = &openapi31.Reflector{}

	// 设置基本信息
	reflector.Spec = &openapi31.Spec{Openapi: "3.1.0"}
	reflector.Spec.Info.
		WithTitle(title).
		WithVersion(version).
		WithDescription(desc)
}

// IsValidReflector 判断 reflector 是否有效
// 当前只前端是否为 nil
func IsValidReflector() bool {
	return reflector != nil
}

// AddRouter 添加一个路由
func AddRouter(path string, method string, input interface{}) {
	if !IsValidReflector() {
		return
	}

	path = parsePath(path)

	oc, err := reflector.NewOperationContext(method, path)
	if err != nil {
		panic(err)
	}
	oc.AddReqStructure(input)

	err = reflector.AddOperation(oc)
	if err != nil {
		panic(err)
	}
}

// parsePath 由于 gin 的路由标识符使用符号， 需要做对应转换
// path parse: "/api/:id" => "/api/{id}"
func parsePath(path string) string {
	// path parse: "/api/:id" => "/api/{id}"
	parts := strings.Split(path, "/")
	for i := range parts {
		part := parts[i]
		if strings.HasPrefix(part, `:`) {
			parts[i] = fmt.Sprintf("{%s}", part[1:])
		}
	}
	path = strings.Join(parts, "/")

	return path
}

// Output 输出 OpenAPI yaml 文件到 os.Stdout
func OutputToStdout() {
	output(os.Stdout)
}

// OutputToFile 输出 OpenAPI Yaml 到文件
func OutputToFile(name string) {
	if !IsValidReflector() {
		return
	}

	f, err := os.OpenFile(name, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	output(f)
}

func output(w io.Writer) {
	if !IsValidReflector() {
		return
	}

	schema, err := reflector.Spec.MarshalYAML()
	if err != nil {
		panic(err)
	}

	_, err = w.Write(schema)
	if err != nil {
		panic(err)
	}
}

// type Option func()

// func WithOptions(opts ...Option) {
// 	for i := range opts {
// 		opts[i]()
// 	}
// }

// func WithTitle(title string) Option {
// 	return func() {
// 		reflector.Spec.Info.WithTitle(title)
// 	}
// }

// func WithVersion(version string) Option {
// 	return func() {
// 		reflector.Spec.Info.WithVersion(version)
// 	}
// }

// func WithDescription(desc string) Option {
// 	return func() {
// 		reflector.Spec.Info.WithDescription(desc)
// 	}
// }
