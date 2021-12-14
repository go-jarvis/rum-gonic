package rum

import (
	"fmt"
	"reflect"
)

func routePath(v interface{}) (string, error) {
	rv := reflect.ValueOf(v)
	rv = reflect.Indirect(rv)

	rt := rv.Type()
	for i := 0; i < rv.NumField(); i++ {

		// 检测 path tag 是否在 httpx.MethoxXXX字段上
		// 判断字段是否具有 Method 方法
		fv := rv.Field(i)
		mfv := fv.MethodByName("Method")
		if !mfv.IsValid() {
			continue
		}

		// 反射类型查找 tag
		ft := rt.Field(i)

		// 反射拿到 methods tag
		if methods, ok := ft.Tag.Lookup("methods"); ok {

			ffv := fv.FieldByName("Methods")
			if ffv.IsValid() && ffv.IsZero() {

				ffv.Set(reflect.ValueOf(methods))
			}
		}

		// 寻找路径
		if val, ok := ft.Tag.Lookup("path"); ok {
			return val, nil
		}

	}
	return "", fmt.Errorf(Err_NoPathProvide)
}
