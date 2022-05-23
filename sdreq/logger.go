package sdreq

import (
	"github.com/gaorx/stardust4/sdlog"
	"github.com/imroc/req/v3"
)

var DefaultLogger req.Logger = logger{}

type logger struct {
}

func (_ logger) Errorf(format string, v ...any) {
	sdlog.Errorf(format, v...)
}

func (_ logger) Warnf(format string, v ...any) {
	sdlog.Warnf(format, v...)
}

func (_ logger) Debugf(format string, v ...any) {
	sdlog.Debugf(format, v...)
}
