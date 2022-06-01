package sdreflect

import (
	"reflect"
)

var (
	TErr = TypeOf[error]()
	TAny = TypeOf[any]()
)

func TypeOf[T any]() reflect.Type {
	return reflect.TypeOf((*T)(nil)).Elem()
}
