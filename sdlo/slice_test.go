package sdlo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSlice(t *testing.T) {
	assert.Nil(t, SliceCopy[int](nil))

	a := []int{1, 2, 3}
	assert.True(t, SliceEqual(a, SliceCopy(a)))

	a = []int{}
	assert.True(t, SliceEqual(a, SliceCopy(a)))

	a = []int{3, 2, 1}
	assert.True(t, SliceEqual(a, a))
	assert.True(t, SliceEqual[int](nil, nil))
	assert.False(t, SliceEqual(a, nil))
	assert.False(t, SliceEqual(nil, a))
	assert.True(t, SliceEqual(a, SliceCopy(a)))
}
