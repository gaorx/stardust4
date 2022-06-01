package sdmansys

import (
	"net/http"
	"strings"

	"github.com/gaorx/stardust4/sdecho"
	"github.com/gaorx/stardust4/sderr"
	"github.com/gaorx/stardust4/sdurl"
	"github.com/labstack/echo/v4"
)

func (sys *System) NewEcho(state any) (*echo.Echo, error) {
	var raws rawEndpoints

	// add endpoints
	for _, ep := range sys.Endpoints {
		rawPath := sdurl.JoinPath(ep.Path)
		if raws.has(rawPath) {
			return nil, sderr.Newf("sdmansys duplicated endpoint '%s'", ep.Path)
		}
		rawMethods, ok := checkHttpMethod(ep.HttpMethods)
		if !ok {
			return nil, sderr.Newf("sdmansys illegal methods '%s'", strings.Join(ep.HttpMethods, "|"))
		}
		rawHandler, ok := ep.toRawHandler()
		if !ok {
			return nil, sderr.Newf("sdmansys illegal endpoint handler '%s'", ep.Path)
		}
		raws = append(raws, rawEndpoint{
			path:    rawPath,
			methods: rawMethods,
			handler: rawHandler,
		})
	}

	// add methods
	for _, m := range sys.Methods {
		rawPath := sdurl.JoinPath(m.PathPrefix, m.Category, m.Name)
		if raws.has(rawPath) {
			return nil, sderr.Newf("sdmansys duplicated method '%s/%s'", m.Category, m.Name)
		}
		rawHandler, ok := m.toRawHandler(sys.ResultTemplate)
		if !ok {
			return nil, sderr.Newf("sdmansys illegal method handler '%s/%s'", m.Category, m.Name)
		}
		raws = append(raws, rawEndpoint{
			path:    rawPath,
			methods: []string{http.MethodPost},
			handler: rawHandler,
		})
	}

	app := sdecho.New(sdecho.Options{})
	for _, raw := range raws {
		if raw.methodIsAny() {
			app.Any(raw.path, raw.handler.toEchoHandler(state))
		} else {
			for _, m := range raw.methods {
				app.Add(m, raw.path, raw.handler.toEchoHandler(state))
			}
		}
	}
	return app, nil
}

func (sys *System) MustEcho(state any) *echo.Echo {
	app, err := sys.NewEcho(state)
	if err != nil {
		panic(err)
	}
	return app
}

type rawHandler func(any, echo.Context) error

type rawEndpoint struct {
	path    string
	methods []string
	handler rawHandler
}

type rawEndpoints []rawEndpoint

func (h rawHandler) toEchoHandler(state any) func(echo.Context) error {
	return func(c echo.Context) error {
		return h(state, c)
	}
}

func (raw rawEndpoint) methodIsAny() bool {
	return len(raw.methods) <= 0
}

func (raws rawEndpoints) has(path string) bool {
	for _, raw := range raws {
		if raw.path == path {
			return true
		}
	}
	return false
}
