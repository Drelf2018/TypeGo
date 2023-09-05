package Pool

import (
	"errors"
	"reflect"
	"sync"
)

type Var interface {
	Set(...any)
	Reset()
}

type None struct{}

func (*None) Set(...any) {}
func (*None) Reset()     {}

type TypePool[T Var] struct {
	sync.Pool
}

func (p *TypePool[T]) Get(args ...any) T {
	t := p.Pool.Get().(T)
	t.Set(args...)
	return t
}

func (p *TypePool[T]) Put(args ...T) {
	for _, arg := range args {
		arg.Reset()
		p.Pool.Put(arg)
	}
}

var ErrType = errors.New("TypeError")

func New[F any, T Var](t T) TypePool[T] {
	if reflect.TypeOf(t) != reflect.TypeOf(new(F)) {
		panic(ErrType)
	}
	return TypePool[T]{sync.Pool{New: func() any { return new(F) }}}
}
