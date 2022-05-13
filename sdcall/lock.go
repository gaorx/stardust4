package sdcall

import (
	"github.com/gaorx/stardust4/sdsync"
)

var (
	Lock  = sdsync.Lock
	LockR = sdsync.LockR
	LockW = sdsync.LockW
)
