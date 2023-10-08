package openapi31

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/swaggest/openapi-go/openapi31"
)

type reflector struct {
	refl *openapi31.Reflector

	writer io.Writer // 输出对象
	file   string    // filename 保存到文件

}

var r *reflector

func Init() *reflector {
	r = &reflector{
		refl:   &openapi31.Reflector{},
		writer: os.Stdout,
	}

	r.refl.Spec = &openapi31.Spec{Openapi: "3.1.0"}
	r.refl.Spec.Info.
		WithTitle("title").
		WithVersion("v0.0.0").
		WithDescription("app description")

	return r
}

// write 将 openapi yaml 格式内容输出到 writer 中。
func (r *reflector) write() {
	if len(r.file) != 0 {
		f, err := os.OpenFile(r.file, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		r.writer = f
	}

	schema, err := r.refl.Spec.MarshalYAML()
	if err != nil {
		panic(err)
	}

	_, err = r.writer.Write(schema)
	if err != nil {
		panic(err)
	}
}

type Option = func(r *reflector)

// WithOptions 添加 options
func WithOptions(opts ...Option) {
	if !IsValid() {
		return
	}

	for _, opt := range opts {
		opt(r)
	}
}

// WithFile 设置保存 openapi 数据的文件。 如果不设置， 将输出到 os.Stdout
func WithFile(output string) Option {
	return func(r *reflector) {
		r.file = output
	}
}

// WithTitle 设置 openapi info: app title
func WithTitle(title string) Option {
	return func(r *reflector) {
		r.refl.Spec.Info.WithTitle(title)
	}
}

// WithVersion 设置 openapi info: app version
func WithVersion(version string) Option {
	return func(r *reflector) {
		r.refl.Spec.Info.WithVersion(version)
	}
}

// WithDescription 设置 openapi info: app description
func WithDescription(dest string) Option {
	return func(r *reflector) {
		r.refl.Spec.Info.WithDescription(dest)
	}
}

// IsValid 判断 reflector 是否有效
// 当前只前端是否为 nil
func IsValid() bool {
	return r != nil
}

// AddRouter 添加一个路由
func AddRouter(path string, method string, input interface{}) {
	if !IsValid() {
		return
	}

	path = parsePath(path)

	oc, err := r.refl.NewOperationContext(method, path)
	if err != nil {
		panic(err)
	}
	oc.AddReqStructure(input)

	err = r.refl.AddOperation(oc)
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

func Output() {
	if !IsValid() {
		return
	}

	r.write()
}
