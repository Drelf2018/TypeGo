package Pool

import (
	"reflect"
	"sync"
)

type Var interface {
	New()
	Set(...any)
	Reset()
}

type Pool[T Var] struct {
	sync.Pool
}

func (p *Pool[T]) Get(args ...any) T {
	t := p.Pool.Get().(T)
	t.Set(args...)
	return t
}

func (p *Pool[T]) Put(ts ...T) {
	for _, t := range ts {
		t.Reset()
		p.Pool.Put(t)
	}
}

func zero(typ reflect.Type) any {
	if typ.Kind() == reflect.Ptr {
		return reflect.New(typ.Elem()).Interface()
	}
	return reflect.Zero(typ).Interface()
}

func New[T Var](t T) (p Pool[T]) {
	typ := reflect.TypeOf(t)
	p.New = func() any {
		i := zero(typ).(T)
		i.New()
		return i
	}
	return
}

// None is a null struct which implement Var(interface).
//
// You can composite it in your own struct so that you don't need to implement Var.
type None struct{}

func (n *None) New()       {}
func (n *None) Set(...any) {}
func (*None) Reset()       {}
