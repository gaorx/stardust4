package sdmansys

import (
	"net/http"
	"reflect"
)

func checkHttpMethod(methods []string) ([]string, bool) {
	if len(methods) <= 0 {
		return nil, true
	}
	for _, m := range methods {
		if m != http.MethodGet &&
			m != http.MethodPost &&
			m != "ANY" &&
			m != "*" &&
			m != http.MethodPut &&
			m != http.MethodDelete &&
			m != http.MethodPatch &&
			m != http.MethodHead &&
			m != http.MethodTrace &&
			m != http.MethodConnect &&
			m != http.MethodOptions {
			return nil, false
		}
		if m == "ANY" || m == "*" {
			return nil, true
		}
	}
	return methods, true
}

func getErr(v reflect.Value) error {
	if !v.IsValid() {
		return nil
	}
	if v.IsNil() {
		return nil
	}
	return v.Interface().(error)
}
