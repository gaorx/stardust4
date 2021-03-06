package sdcall

import (
	"sync"
)

func Parallel(concurrent int, actions []func()) error {
	numActions := len(actions)
	if numActions == 0 {
		return nil
	}
	if concurrent <= 0 {
		var wg sync.WaitGroup
		for _, f := range actions {
			wg.Add(1)
			go func(f func()) {
				defer wg.Done()
				_ = Safe(f)
			}(f)
		}
		wg.Wait()
		return nil
	} else {
		if concurrent > numActions {
			concurrent = numActions
		}
		pool, err := NewPool(concurrent, &PoolOptions{
			PreAlloc: true,
		})
		if err != nil {
			return err
		}
		defer func() { _ = pool.Close() }()
		var wg sync.WaitGroup
		for _, f := range actions {
			f1 := f
			wg.Add(1)
			err := pool.Submit(func() {
				defer wg.Done()
				_ = Safe(f1)
			})
			if err != nil {
				return err
			}
		}
		wg.Wait()
		return nil
	}
}

func ParallelSlice[T any](concurrency int, l []T, action func(int, T)) error {
	return Parallel(concurrency, Bind(l, action))
}
