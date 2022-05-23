package sdsemver

import (
	"testing"

	"github.com/gaorx/stardust4/sdrand"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	sdrand.InitSeed()
	_, err := Parse("")
	assert.Error(t, err)
	_, err = Parse("a.b.c")
	assert.Error(t, err)
	_, err = Parse("1000000.1.1")
	assert.Error(t, err)
	_, err = Parse("1.10000000.1")
	assert.Error(t, err)
	_, err = Parse("1.0.1000000")
	assert.Error(t, err)
	v, err := Parse("3")
	assert.True(t, v.Equal(3, 0, 0))
	v, err = Parse("0.3")
	assert.True(t, v.Equal(0, 3, 0))
	v, err = Parse("0.2.3")
	assert.True(t, v.Equal(0, 2, 3))
	_, err = Parse("0.2.3.4")
	assert.Error(t, err)
}
