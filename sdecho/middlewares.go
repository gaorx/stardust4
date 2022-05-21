package sdecho

import (
	"fmt"
	"time"

	"github.com/gaorx/stardust4/sdcall"
	"github.com/gaorx/stardust4/sderr"
	"github.com/gaorx/stardust4/sdlog"
	"github.com/gaorx/stardust4/sdparse"
	"github.com/gaorx/stardust4/sdtime"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LoggingRecover(logSkipper middleware.Skipper) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ec echo.Context) error {
			if logSkipper != nil && logSkipper(ec) {
				return next(ec)
			}
			req := ec.Request()
			res := ec.Response()
			startAt := time.Now()
			var nextErr, panicErr, finalErr error
			panicErr = sdcall.Safe(func() {
				nextErr = next(ec)
			})
			if panicErr != nil {
				finalErr = panicErr
			} else {
				finalErr = nextErr
			}
			if finalErr != nil {
				ec.Error(finalErr)
			}
			elapsedHuman := time.Since(startAt)
			elapsedMs := sdtime.ToMillisF(elapsedHuman)
			statusCode := res.Status
			method := req.Method
			path := req.URL.Path
			if path == "" {
				path = "/"
			}

			bytesIn, err := sdparse.Int64(req.Header.Get(echo.HeaderContentLength))
			if err != nil {
				bytesIn = 0
			}

			logFields := sdlog.Fields{
				"latency":   elapsedMs,
				"latency_h": elapsedHuman,
				"remote_ip": ec.RealIP(),
				"bytes_in":  bytesIn,
				"bytes_out": res.Size,
			}
			if finalErr == nil {
				sdlog.WithFields(logFields).Infof("%d %s %s", statusCode, method, path)
			} else {
				logFields["error"] = fmt.Sprintf("%+v", finalErr)
				sdlog.WithFields(logFields).Infof("%d %s %s", statusCode, method, path)
			}
			return sderr.Wrap(finalErr, "sdecho logging recover middleware error")
		}
	}
}
