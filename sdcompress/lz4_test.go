package sdcompress

import (
	"testing"

	"github.com/gaorx/stardust4/sdbytes"
	"github.com/gaorx/stardust4/sdrand"
	"github.com/stretchr/testify/assert"
)

func TestLz4(t *testing.T) {
	sdrand.InitSeed()
	data0 := []byte(sdrand.String(1303, sdrand.LowerLetterNumbers))
	for _, level := range Lz4AllLevels {
		data1, err := Lz4(data0, level)
		assert.NoError(t, err)
		data2, err := Unlz4(data1)
		assert.NoError(t, err)
		assert.True(t, sdbytes.Equal(data0, data2))
	}
}
