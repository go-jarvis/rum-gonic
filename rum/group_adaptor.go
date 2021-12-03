package rum

// Use 添加中间件， 兼容 gin
func (r *RouterGroup) Use(fns ...HandlerFunc) {
	for _, fn := range fns {
		mid := NewMiddleware(fn)
		r.operators = append(r.operators, mid)
	}
}

func (r *RouterGroup) StatisFile() {}
