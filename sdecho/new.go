package sdecho

import (
	"net/http"
	"strings"

	"github.com/gaorx/stardust4/sdlog"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Options struct {
	DebugMode    bool
	LogSkipper   middleware.Skipper
	ErrorHandler echo.HTTPErrorHandler
}

func New(opts Options) *echo.Echo {
	if opts.ErrorHandler == nil {
		opts.ErrorHandler = defaultHttpErrorHandler
	}
	app := echo.New()
	app.Debug = opts.DebugMode
	app.HideBanner = true
	app.HidePort = true
	app.Use(LoggingRecover(opts.LogSkipper))
	app.HTTPErrorHandler = opts.ErrorHandler
	return app
}

func defaultHttpErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	he, ok := err.(*echo.HTTPError)
	if ok {
		if he.Internal != nil {
			if herr, ok := he.Internal.(*echo.HTTPError); ok {
				he = herr
			}
		}
	} else {
		errMsg := http.StatusText(http.StatusInternalServerError)
		if c.QueryParam("_show_error") == "1" {
			errMsg = errMsg + strings.Repeat("\r\n", 2) + err.Error()
		}
		he = &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: errMsg,
		}
	}

	var errMsg string
	if m, ok := he.Message.(string); ok {
		errMsg = m
	} else {
		errMsg = "Unknown error"
	}

	if c.Request().Method == http.MethodHead {
		err = c.NoContent(he.Code)
	} else {
		err = c.String(he.Code, errMsg)
	}
	if err != nil {
		sdlog.Errorf("http error handler error: %s", err)
	}
}
