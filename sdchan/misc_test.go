package sdchan

import (
	"testing"

	"github.com/gaorx/stardust4/sdrand"
	"github.com/gaorx/stardust4/sdtime"
	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	sdrand.InitSeed()
	total := 400

	makeChan := func(odd bool) (chan int, func()) {
		c := make(chan int)
		return c, func() {
			if odd {
				for i := 0; i < total; i++ {
					if i%2 == 1 {
						c <- i
						sdtime.SleepMs(sdrand.Int64r(2, 7))
					}
				}
			} else {
				for i := 0; i < total; i++ {
					if i%2 == 0 {
						c <- i
						sdtime.SleepMs(sdrand.Int64r(4, 10))
					}
				}
			}
			close(c)
		}
	}

	oddChan, oddStarter := makeChan(true)
	evenChan, evenStarter := makeChan(false)
	merged := MergeRecv(oddChan, evenChan)
	go func() { oddStarter() }()
	go func() { evenStarter() }()

	counter := map[int]int{}
	for {
		v, ok := <-merged
		if !ok {
			break
		}
		counter[v] = counter[v] + 1
	}
	for i := 0; i < total; i++ {
		assert.Equal(t, 1, counter[i])
	}
}
