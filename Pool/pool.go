package Pool

import (
	"sync"
)

type Var[T any] interface {
	OnNew()
	OnGet(...any)
	OnPut()
	~*T
}

type Pool[T any, PT Var[T]] struct {
	len int
	cap int
	sync.Pool
}

func (p *Pool[T, PT]) Get(args ...any) (pt PT) {
	pt = p.Pool.Get().(PT)
	pt.OnGet(args...)
	p.len++
	return
}

func (p *Pool[T, PT]) Put(pt ...PT) {
	for _, v := range pt {
		if v == nil {
			continue
		}
		v.OnPut()
		p.Pool.Put(v)
		p.len--
	}
}

func (p *Pool[T, PT]) Len() int {
	return p.len
}

func (p *Pool[T, PT]) Cap() int {
	return p.cap
}

func New[T any, PT Var[T]](_ ...PT) *Pool[T, PT] {
	p := new(Pool[T, PT])
	p.New = func() any {
		var v PT = new(T)
		v.OnNew()
		p.cap++
		return v
	}
	return p
}
