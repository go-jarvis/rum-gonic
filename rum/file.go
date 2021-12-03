package rum

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FileOperator interface {
	PathOperator
	MethodOperator
	Output(c *gin.Context)
}

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

func (r *RouterGroup) RegisterFile(path, filepath string) {
	for _, method := range []string{http.MethodGet, http.MethodHead} {
		op := NewStaticFile(method, path, filepath)

		r.addOperators(op)
	}
}
