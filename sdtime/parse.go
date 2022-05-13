package sdtime

import (
	"github.com/gaorx/stardust4/sdparse"
)

var (
	Parse     = sdparse.Time
	ParseDef  = sdparse.TimeDef
	MustParse = sdparse.MustTime
)
