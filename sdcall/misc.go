package sdcall

import (
	"github.com/gaorx/stardust4/sderr"
)

func Retry(maxRetries int, action func() error) error {
	if action == nil {
		return nil
	}
	err0 := action()
	if err0 == nil {
		return nil
	}
	err := sderr.Append(err0)
	for i := 1; i <= maxRetries; i++ {
		err0 := action()
		if err0 != nil {
			err = sderr.Append(err, err0)
		} else {
			return nil
		}
	}
	return err
}

func Safe(action func()) (err error) {
	if action == nil {
		err = nil
		return
	}

	defer func() {
		if err0 := recover(); err0 != nil {
			err = sderr.ToErr(err0)
		}
	}()
	action()
	return
}

func Bind[T any](l []T, action func(int, T)) []func() {
	if action == nil {
		action = func(int, T) {}
	}
	actions := make([]func(), 0, len(l))
	for i, v := range l {
		i0, v0 := i, v
		actions = append(actions, func() {
			action(i0, v0)
		})
	}
	return actions
}
