package rum

import (
	"reflect"
)

func routePath(v interface{}) string {
	rv := reflect.Indirect(reflect.ValueOf(v))

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
		if val, ok := ft.Tag.Lookup("path"); ok {
			return val
		}

	}
	return ""
}
