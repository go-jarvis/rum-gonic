package rum

import (
	"context"
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
)

const (
	ctxInjectorPrefix = "rum-with-context"
)

var contextInjectors []ContextInjectorFunc

type ContextInjectorFunc func(ctx context.Context) context.Context

// WithContextValue 将 key, value 注入到 context 中。 同 context.WithValue(ctx, key, value)
// 由于 gin 中 context.Keys = map[string]interface{} , 只能使用 string 作为 key， 因此在保存先进行转换成带类型的值， 降低冲突。
func WithContextValue(key, value interface{}) ContextInjectorFunc {

	_key := keyAsString(key)

	return func(ctx context.Context) context.Context {
		if c, ok := ctx.(*gin.Context); ok {
			c.Set(_key, value)
			return c
		}

		return context.WithValue(ctx, _key, value)
	}
}

// FromContextValue 从 context 中获取对应的值。 同 context.Value(key)。
func FromContextValue(ctx context.Context, key interface{}) interface{} {
	_key := keyAsString(key)
	return ctx.Value(_key)
}

// keyAsString 将对象转换为字符串。 v 需要是 string 或者实现 fmt.Stringer 接口，即 String() 方法。
func keyAsString(v interface{}) string {
	return fmt.Sprintf("%s-%s-%s", ctxInjectorPrefix, typeName(v), stringify(v))
}

// deReflectType 获取对象的真实类型
func deReflectType(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return t
}

// stringify 返回类型的值
func stringify(v interface{}) string {
	switch s := v.(type) {
	case fmt.Stringer:
		return s.String()
	case string:
		return s
	}

	return fmt.Sprint(v)
	// panic("<not Stringer>")
	// return "<not Stringer>"
}

// typeName 返回类型的名称
func typeName(v interface{}) string {
	rt := reflect.TypeOf(v)
	rt = deReflectType(rt)

	return rt.Name()
}
