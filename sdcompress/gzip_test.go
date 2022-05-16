package sdcompress

import (
	"testing"

	"github.com/gaorx/stardust4/sdbytes"
	"github.com/gaorx/stardust4/sdrand"
	"github.com/stretchr/testify/assert"
)

func TestGzip(t *testing.T) {
	sdrand.InitSeed()
	data0 := []byte(sdrand.String(1303, sdrand.LowerLetterNumbers))
	for _, level := range GzipAllLevels {
		data1, err := Gzip(data0, level)
		assert.NoError(t, err)
		data2, err := Ungzip(data1)
		assert.NoError(t, err)
		assert.True(t, sdbytes.Equal(data0, data2))
	}
}
