package sdtext

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TextSimplifiedChinese(t *testing.T) {
	testEncoding := func(e CharsetEncoding, s string) {
		encoded, err := e.Encode(s)
		assert.NoError(t, err)
		decoded, err := e.Decode(encoded)
		assert.NoError(t, err)
		assert.True(t, s == decoded)
	}

	testEncoding(GBK, "你好，世界")
	testEncoding(GB2312, "你好，世界")
	testEncoding(GB18030, "你好，世界")
}
