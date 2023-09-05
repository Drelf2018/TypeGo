package Pool

import (
	"reflect"
	"sync"
)

type pool[T any] struct {
	reset func(T)
	inner *sync.Pool
}

func (p *pool[T]) Get() T {
	return p.inner.Get().(T)
}

func (p *pool[T]) Put(args ...T) {
	for _, arg := range args {
		p.reset(arg)
		p.inner.Put(arg)
	}
}

func (p *pool[T]) Set(reset func(T)) {
	if reset != nil {
		p.reset = reset
	}
}

var r = reflect.TypeOf((*interface{ Reset() })(nil)).Elem()

func New[T any](t *T) pool[*T] {
	p := pool[*T]{
		reset: func(t *T) {},
		inner: &sync.Pool{New: func() any { return new(T) }},
	}
	if typ := reflect.TypeOf(t); typ.Implements(r) {
		if f, ok := typ.MethodByName(r.Method(0).Name); ok {
			p.reset = func(t *T) { f.Func.Call([]reflect.Value{reflect.ValueOf(t)}) }
		}
	}
	return p
}
