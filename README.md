# rum-gonic

todo list: 

逻辑路由
    + 使用 **反射获取 path 路径**
```go
type Index struct {
    httpx.MethodGet `path:"/index"`
}

func (index *Index) Output(c *gin.Context) (interface{},error){}


r.Register(&Index{})
```

中间件

```go
func NewMiddleware(fn HandlerFunc) *Middleware {
	return &Middleware{
		middwareFunc: fn,
	}
}

r.Register(mid)
```