package sdcall

import (
	"sync"

	"github.com/gaorx/stardust4/sderr"
	"github.com/panjf2000/ants/v2"
)

type Pool struct {
	pool *ants.Pool
}

type PoolOptions = ants.Options

var (
	ErrInvalidPoolExpiry   = ants.ErrInvalidPoolExpiry
	ErrLackPoolFunc        = ants.ErrLackPoolFunc
	ErrPoolClosed          = ants.ErrPoolClosed
	ErrPoolOverload        = ants.ErrPoolOverload
	ErrInvalidPreAllocSize = ants.ErrInvalidPreAllocSize
)

func NewPool(size int, opts *PoolOptions) (*Pool, error) {
	var antsOpts []ants.Option
	if opts != nil {
		antsOpts = append(antsOpts, ants.WithOptions(*opts))
	}
	p, err := ants.NewPool(size, antsOpts...)
	if err != nil {
		return nil, sderr.Wrap(err, "sdcall create ants pool error")
	}
	return &Pool{pool: p}, nil
}

func (p *Pool) NumFree() int {
	return p.pool.Free()
}

func (p *Pool) NumCap() int {
	return p.pool.Cap()
}

func (p *Pool) NumRunning() int {
	return p.pool.Running()
}

func (p *Pool) Close() error {
	p.pool.Release()
	return nil
}

func (p *Pool) Submit(action func()) error {
	if action == nil {
		return nil
	}
	err := p.pool.Submit(action)
	return sderr.Wrap(err, "sdcall submit action error")
}

func (p *Pool) Do(action func()) error {
	if action == nil {
		return nil
	}
	var wg sync.WaitGroup
	wg.Add(1)
	err := p.pool.Submit(func() {
		defer wg.Done()
		_ = Safe(action)
	})
	if err != nil {
		return err
	}
	wg.Wait()
	return nil
}

func (p *Pool) Wrap(action func()) func() {
	if action == nil {
		return nil
	}
	return func() {
		_ = p.Do(action)
	}
}
