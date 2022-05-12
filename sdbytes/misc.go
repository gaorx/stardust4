package sdbytes

import (
	"github.com/gaorx/stardust4/sdlo"
)

var (
	Copy  = sdlo.SliceCopy[byte]
	Equal = sdlo.SliceEqual[byte]
)
