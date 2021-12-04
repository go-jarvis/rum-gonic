package rum

import "reflect"

const (
	withContextPrefix = "rum-with-context-"
)

type WithContextOperator interface {
	String()
}

var withContextOperators map[string]interface{}

// func WithContext(key, value interface{}) {
// 	sKey := WithContextKey(key)

// 	withContext(sKey, value)
// }

func withContext(key string, value interface{}) {
	if withContextOperators == nil {
		withContextOperators = make(map[string]interface{})
	}

	// sKey := WithContextKey(key)
	if _, ok := withContextOperators[key]; ok {
		panic("key exists in withContextOperator")
	}

	withContextOperators[key] = value

}

func WithContextKey(key interface{}) string {
	return withContextPrefix + typeName(key)
}

func typeName(v interface{}) string {
	rt := reflect.TypeOf(v)

	return deref(rt).Name()
}

func deref(rt reflect.Type) reflect.Type {
	for rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	return rt
}

type ContextInjection struct {
	key   string
	value interface{}
}

func NewContextInjector(key interface{}, value interface{}) *ContextInjection {
	_key := WithContextKey(key)
	return &ContextInjection{
		key:   _key,
		value: value,
	}
}

func (injector *ContextInjection) Key() string {
	return injector.key
}

func (injector *ContextInjection) Value() interface{} {
	return injector.value
}

type ContextInjector interface {
	Key() string
	Value() interface{}
}
