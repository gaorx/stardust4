package sdmansys

import (
	"reflect"

	"github.com/gaorx/stardust4/sdecho/sdhttpapi"
	"github.com/gaorx/stardust4/sderr"
	"github.com/gaorx/stardust4/sdreflect"
	"github.com/labstack/echo/v4"
)

type Method struct {
	PathPrefix string
	Category   string
	Name       string
	Facade     any
	Handler    any
}

func (sys *System) AddMethod(category, name string, h any, facade any) {
	sys.Methods = append(sys.Methods, Method{
		PathPrefix: sys.PathPrefix,
		Category:   category,
		Name:       name,
		Facade:     facade,
		Handler:    h,
	})
}

func (sys *System) AddMethods(category string, handlers []Method, facade any) {
	for _, h := range handlers {
		if h.Category == "" {
			h.Category = category
		}
		if h.Facade == nil {
			h.Facade = facade
		}
		sys.AddMethod(h.Category, h.Name, h.Facade, h.Handler)
	}
}

func (m Method) toRawHandler(rt *sdhttpapi.ResultTemplate) (rawHandler, bool) {
	if m.Handler == nil {
		return nil, false
	}
	if rt == nil {
		rt = sdhttpapi.DefaultResultTemplate()
	}
	facade := m.Facade
	hv := sdreflect.ValueOf(m.Handler)
	ht := hv.Type()
	if ht.Kind() != reflect.Func {
		return nil, false
	}

	var inTyp1, inTyp2 reflect.Type
	switch ht.NumIn() {
	case 0:
		inTyp1, inTyp2 = nil, nil
	case 1:
		inTyp1, inTyp2 = ht.In(0), nil
	case 2:
		inTyp1, inTyp2 = ht.In(0), ht.In(1)
	default:
		return nil, false
	}

	var outWithErr bool
	switch ht.NumOut() {
	case 0:
		outWithErr = false
	case 1:
		outWithErr = ht.Out(0) == sdreflect.TErr
	case 2:
		if ht.Out(1) != sdreflect.TErr {
			return nil, false
		}
		outWithErr = true
	default:
		return nil, false
	}

	return func(state any, c echo.Context) error {
		var inVals []reflect.Value
		if inTyp1 != nil && inTyp2 != nil {
			sv := sdreflect.ValueOf(state)
			inVal2 := reflect.New(inTyp2)
			if err := c.Bind(inVal2.Interface()); err != nil {
				return rt.Err(sderr.Wrap(err, "sdmansys bind request json error")).Render(c)
			}
			inVals = append(inVals, sv, inVal2.Elem())
		} else if inTyp1 != nil && inTyp2 == nil {
			sv := sdreflect.ValueOf(state)
			if state != nil && sv.Type() == inTyp1 {
				inVals = append(inVals, sv)
			} else {
				inVal1 := reflect.New(inTyp1)
				if err := c.Bind(inVal1.Interface()); err != nil {
					return rt.Err(sderr.Wrap(err, "sdmansys bind request json error")).Render(c)
				}
				inVals = append(inVals, inVal1.Elem())
			}
		}

		outVals := hv.Call(inVals)

		switch len(outVals) {
		case 0:
			return rt.Ok(nil).WithFacade(facade).Render(c)
		case 1:
			if outWithErr {
				callErr := getErr(outVals[0])
				return rt.Err(sderr.Wrap(callErr, "sdmansys call method error")).Render(c)
			} else {
				callOut := outVals[0].Interface()
				return rt.Ok(callOut).WithFacade(facade).Render(c)
			}
		case 2:
			callOut, callErr := outVals[0].Interface(), getErr(outVals[1])
			return rt.Of(callOut, callErr).WithFacade(facade).Render(c)
		default:
			return rt.Err(sderr.New("sdmansys run here error")).Render(c)
		}
	}, true
}
