package sdmansys

import (
	"github.com/gaorx/stardust4/sdecho"
	"github.com/gaorx/stardust4/sdstrings"
	"github.com/labstack/echo/v4"
)

type Endpoint struct {
	Path        string
	HttpMethods []string
	Handler     any
}

func (sys *System) AddEndpoint(method, path string, h any) *System {
	sys.Endpoints = append(sys.Endpoints, Endpoint{
		Path:        path,
		HttpMethods: sdstrings.SplitNonempty(method, "|", true),
		Handler:     h,
	})
	return sys
}

func (ep Endpoint) toRawHandler() (rawHandler, bool) {
	switch h1 := ep.Handler.(type) {
	case nil:
		return nil, false
	case func(any, echo.Context) error:
		return h1, true
	case func(any, sdecho.Context) error:
		return func(state any, c echo.Context) error {
			return h1(state, sdecho.C(c))
		}, true
	case func(c echo.Context) error:
		return func(_ any, c echo.Context) error {
			return h1(c)
		}, true
	case func(c sdecho.Context) error:
		return func(_ any, c echo.Context) error {
			return h1(sdecho.C(c))
		}, true
	default:
		panic("TODO")
	}
}
