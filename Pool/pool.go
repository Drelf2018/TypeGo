package Pool

import "sync"

type Var interface {
	New()
	Set(...any)
	Reset()
}

type Pointer[S any] interface {
	*S
	Var
}

type TypePool[T Var] struct {
	sync.Pool
}

func (p *TypePool[T]) Get(args ...any) T {
	t := p.Pool.Get().(T)
	t.Set(args...)
	return t
}

func (p *TypePool[T]) Put(ts ...T) {
	for _, t := range ts {
		t.Reset()
		p.Pool.Put(t)
	}
}

func New[S any, P Pointer[S]](_ P) TypePool[P] {
	return TypePool[P]{
		sync.Pool{
			New: func() any {
				t := new(S)
				P.New(t)
				return t
			},
		},
	}
}

// None is a null struct which implement Var(interface).
//
// You can composite it in your own struct so that you don't need to implement Var.
type None struct{}

func (n *None) New()       {}
func (n *None) Set(...any) {}
func (*None) Reset()       {}
