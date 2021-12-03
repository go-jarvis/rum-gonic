package rum

import "reflect"

func routePath(v interface{}) string {
	rv := reflect.Indirect(reflect.ValueOf(v))

	rt := rv.Type()
	for i := 0; i < rv.NumField(); i++ {
		ft := rt.Field(i)
		if val, ok := ft.Tag.Lookup("path"); ok {
			return val
		}

	}
	return ""
}
