package sdcall

import (
	"testing"

	"github.com/gaorx/stardust4/sderr"
	"github.com/stretchr/testify/assert"
)

func TestRetry(t *testing.T) {
	err0 := sderr.New("0")

	var n int
	fail := func() error {
		n++
		return err0
	}
	success := func() error {
		n++
		return nil
	}

	{
		n = 0
		err := Retry(0, success)
		assert.NoError(t, err)
		assert.Equal(t, 1, n)
	}
	{
		n = 0
		err := Retry(0, fail)
		assert.True(t, sderr.Is(err, err0))
		assert.Equal(t, 1, n)
	}

	{
		n = 0
		err := Retry(2, success)
		assert.NoError(t, err)
		assert.Equal(t, 1, n)
	}

	{
		n = 0
		err := Retry(2, fail)
		assert.True(t, sderr.Is(err, err0))
		assert.Equal(t, 3, n)
	}
}

func TestSafe(t *testing.T) {
	err0 := sderr.New("err0")

	fail := func() { panic(err0) }
	success := func() {}

	{
		err := Safe(success)
		assert.NoError(t, err)
	}

	{
		err := Safe(fail)
		assert.True(t, sderr.Is(err, err0))
	}
}
