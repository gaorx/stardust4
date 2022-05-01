package sderr

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAndAs(t *testing.T) {
	root1 := myErr1{1}
	root2 := &myErr2{2}
	err1 := Wrap(root1, "wrap1")
	err2 := Wrap(root2, "wrap2")

	// is
	assert.True(t, Is(err1, root1))
	assert.False(t, Is(err1, root2))
	assert.True(t, Is(err2, root2))
	assert.False(t, Is(err2, root1))

	// as (myErr1)
	root1a, ok := As[myErr1](err1)
	assert.True(t, ok)
	assert.Equal(t, root1, root1a)
	root1b, ok := As[*myErr2](err1)
	assert.False(t, ok)
	assert.Nil(t, root1b)
	// as (*myErr2)
	root2a, ok := As[*myErr2](err2)
	assert.True(t, ok)
	assert.Equal(t, root2, root2a)
	root2b, ok := As[myErr1](err2)
	assert.False(t, ok)
	assert.Equal(t, myErr1{}, root2b)
}

type myErr1 struct {
	Code int
}

func (e myErr1) Error() string {
	return fmt.Sprintf("myErr1(%d)", e.Code)
}

type myErr2 struct {
	Code int
}

func (e *myErr2) Error() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf("myErr2(%d)", e.Code)
}
