package sdreflect

import (
	"reflect"
)

func ValueOf(v any) reflect.Value {
	if v == nil {
		return reflect.Value{}
	}
	switch v1 := v.(type) {
	case reflect.Value:
		return v1
	case *reflect.Value:
		if v1 == nil {
			return reflect.Value{}
		} else {
			return *v1
		}
	default:
		return reflect.ValueOf(v1)
	}
}

func InterfaceOf(v any) any {
	if v == nil {
		return nil
	}
	switch v1 := v.(type) {
	case reflect.Value:
		return v1.Interface()
	case *reflect.Value:
		if v1 == nil {
			return nil
		} else {
			return v1.Interface()
		}
	default:
		return v1
	}
}
