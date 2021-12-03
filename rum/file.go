package rum

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type StaticFile struct {
	method   string
	path     string
	filepath string
}

func (static *StaticFile) Path() string {
	return static.path
}

func (static *StaticFile) Method() string {
	return static.method
}

func (static *StaticFile) Output(c *gin.Context) (interface{}, error) {
	c.File(static.filepath)
	c.Abort()

	return nil, nil
}

func NewStaticFile(method, path, filepath string) *StaticFile {
	return &StaticFile{
		method:   method,
		path:     path,
		filepath: filepath,
	}
}

// StaticFile registers a single route in order to serve a single file of the local filesystem.
// router.StaticFile("favicon.ico", "./resources/favicon.ico")
func (r *RouterGroup) StaticFile(path, filepath string) {
	for _, method := range []string{http.MethodGet, http.MethodHead} {
		op := NewStaticFile(method, path, filepath)

		r.addOperators(op)
	}
}

type StaticFS struct {
	method string
	path   string
	fs     http.FileSystem
}

func (static *StaticFS) Path() string {
	return static.path
}
func (static *StaticFS) Method() string {
	return static.method
}

func (static *StaticFS) Output(c *gin.Context) (interface{}, error) {

	file := c.Param("filepath")
	urlPath := c.Request.URL.Path
	prefix := strings.TrimRight(urlPath, file)

	// https://shockerli.net/post/golang-pkg-http-file-server/#支持子目录路径
	// 在使用 static 或 staticFS 后,  fileserver 的路由工作目录已经切换到文件目录
	// 因此 request url 中是包含了前缀的目录， 需要隐藏
	fileserver := http.StripPrefix(prefix, http.FileServer(static.fs))

	/* Check if file exists and/or if we have permission to access it */
	f, err := static.fs.Open(file)
	if err != nil {
		/* 可行 */
		c.String(http.StatusNotFound, "404 page not found")
		c.Abort()

		return nil, nil
	}
	f.Close()

	fileserver.ServeHTTP(c.Writer, c.Request)

	c.Abort()
	return nil, nil

}

// StaticFS works just like `Static()` but a custom `http.FileSystem` can be used instead.
// Gin by default user: gin.Dir()
func (r *RouterGroup) StaticFS(path string, fs http.FileSystem) {

	if strings.Contains(path, ":") || strings.Contains(path, "*") {
		panic("URL parameters can not be used when serving a static folder")
	}

	for _, method := range []string{http.MethodGet, http.MethodHead} {
		op := &StaticFS{
			method: method,
			path:   fmt.Sprintf("%s/*filepath", path),
			fs:     fs,
		}

		r.addOperators(op)
	}

}

// Static serves files from the given file system root.
// Internally a http.FileServer is used, therefore http.NotFound is used instead
// of the Router's NotFound handler.
// To use the operating system's file system implementation,
// use :
//     router.Static("/static", "/var/www")
func (r *RouterGroup) Static(path string, dirpath string) {
	r.StaticFS(path, gin.Dir(dirpath, false))
}
