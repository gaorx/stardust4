package sdecho

import (
	"github.com/labstack/echo/v4"
)

func H[T any](state T, h func(T, Context) error) echo.HandlerFunc {
	return func(c echo.Context) error {
		return h(state, C(c))
	}
}

func E[T any](state T, h func(T, error, Context)) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		h(state, err, C(c))
	}
}
