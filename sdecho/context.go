package sdecho

import (
	"github.com/labstack/echo/v4"
)

type Context struct {
	echo.Context
}

func C(c echo.Context) Context {
	if c1, ok := c.(Context); ok {
		return c1
	} else {
		return Context{c}
	}
}

func WrapHandler(f func(Context) error) echo.HandlerFunc {
	return func(c echo.Context) error {
		return f(Context{c})
	}
}

func WrapErrorHandler(f func(error, Context)) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		f(err, Context{c})
	}
}

func WrapHandlerWith[T any](state T, f func(T, Context) error) echo.HandlerFunc {
	return func(c echo.Context) error {
		return f(state, Context{c})
	}
}

func WrapErrorHandlerWith[T any](state T, f func(T, error, Context)) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		f(state, err, Context{c})
	}
}
