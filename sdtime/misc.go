package sdtime

import (
	"time"
)

func NowTruncateS() time.Time {
	return time.Now().Truncate(time.Second)
}
