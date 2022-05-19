package sdlo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinMax(t *testing.T) {
	// Min2/Max2
	assert.Equal(t, 4, Max2(3, 4))
	assert.Equal(t, 3, Min2(3, 4))
	assert.Equal(t, 3, Max2(3, 3))
	assert.Equal(t, 3, Min2(3, 3))

	// MinN/MaxN
	assert.Equal(t, 0, MinN[int]())
	assert.Equal(t, 0, MaxN[int]())
	assert.Equal(t, 1, MinN(1))
	assert.Equal(t, 1, MaxN(1))
	assert.Equal(t, 2, MinN(3, 4, 5, 2))
	assert.Equal(t, 5, MaxN(3, 4, 5, 2))
}
