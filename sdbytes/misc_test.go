package sdbytes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMisc(t *testing.T) {
	a := []byte("你好,world")
	b := Copy(a)
	assert.True(t, Equal(a, b))
	assert.True(t, Equal(nil, nil))
	assert.False(t, Equal(nil, []byte{}))
	assert.False(t, Equal([]byte{}, nil))
	assert.True(t, Equal([]byte{}, []byte{}))
}
