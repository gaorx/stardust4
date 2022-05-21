package sdhttpapi

import (
	"reflect"

	"github.com/gaorx/stardust4/sdecho"
	"github.com/gaorx/stardust4/sderr"
	"github.com/gaorx/stardust4/sdjson"
	"github.com/gaorx/stardust4/sdlo"
	"github.com/gaorx/stardust4/sdreflect"
	"github.com/labstack/echo/v4"
)

type Renderer func(c echo.Context, r *Result) error

var (
	ErrFacadeType = sderr.New("sdhttpapi facade type error")
)

func RenderJSON(c echo.Context, r *Result) error {
	c1 := sdecho.Context{Context: c}

	var o sdjson.Object
	if r.IsOk() {
		var facadeData any
		if r.Facade != nil && r.Data != nil {
			dv := reflect.ValueOf(r.Data)
			dt := dv.Type()
			if dt.Kind() == reflect.Slice || dt.Kind() == reflect.Array {
				n := dv.Len()
				facade, ok := normalizeFacade(r.Facade, dt.Elem())
				if !ok {
					return r.Template().Err(sderr.WithStack(ErrFacadeType)).WithRenderer(r.Renderer).Render(c)
				}
				facadeElems := make([]any, 0, n)
				for i := 0; i < n; i++ {
					facadeElem := facade(i, dv.Index(i).Interface())
					facadeElems = append(facadeElems, facadeElem)
				}
				facadeData = facadeElems
			} else {
				facade, ok := normalizeFacade(r.Facade, dt)
				if !ok {
					return r.Template().Err(sderr.WithStack(ErrFacadeType)).WithRenderer(r.Renderer).Render(c)
				}
				facadeData = facade(0, dv.Interface())
			}
		} else {
			facadeData = r.Data
		}

		o = sdjson.Object{"code": r.Code, "data": facadeData}
	} else {
		o = sdjson.Object{"code": r.Code, "error_msg": r.ErrorMessage}
		if c1.Echo().Debug && c1.ArgBool("_error_stack", false) {
			o["error_stack"] = sdlo.SliceNilToEmpty(r.ErrorStack)
		}
	}
	for k, v := range r.Fields {
		if !o.Has(k) {
			o[k] = v
		}
	}

	jsonpCallback := c1.ArgString("_callback", "")
	if jsonpCallback != "" {
		return c.JSONP(200, jsonpCallback, o)
	} else {
		return c.JSON(200, o)
	}
}

func normalizeFacade(facade any, elemTyp reflect.Type) (func(int, any) any, bool) {
	if facade1, ok := facade.(func(any) any); ok {
		return func(_ int, elem any) any {
			return facade1(elem)
		}, true
	}
	if facade1, ok := facade.(func(int, any) any); ok {
		return facade1, true
	}

	df := reflect.ValueOf(facade)
	dft := df.Type()
	if dft.Kind() != reflect.Func {
		return nil, false
	}
	if dft.NumIn() == 1 && dft.NumOut() == 1 {
		if !elemTyp.AssignableTo(dft.In(0)) {
			return nil, false
		}
		return func(_ int, elem any) any {
			drl := df.Call([]reflect.Value{
				sdreflect.ValueOf(elem),
			})
			return drl[0].Interface()
		}, true
	} else if dft.NumIn() == 2 && dft.NumOut() == 1 {
		if !(reflect.TypeOf(0).AssignableTo(dft.In(0)) && elemTyp.AssignableTo(dft.In(1))) {
			return nil, false
		}
		return func(index int, elem any) any {
			drl := df.Call([]reflect.Value{
				reflect.ValueOf(index),
				sdreflect.ValueOf(elem),
			})
			return drl[0].Interface()
		}, true
	} else {
		return nil, false
	}
}
