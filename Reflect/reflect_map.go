package Reflect

import (
	"reflect"
)

type Map[V any] struct {
	Alias
	Parser[V]
	types map[uintptr]V
}

func (r *Map[V]) GetType(elem reflect.Type) (v V, ok bool) {
	// type check
	if elem.Kind() == reflect.Ptr {
		elem = elem.Elem()
	}
	if elem.Kind() != reflect.Struct {
		return
	}
	// exists
	ue := Type(elem)
	if v, ok = r.types[ue]; ok {
		return
	}
	// parse
	r.types[ue] = v
	v, ok = r.Parse(r, elem), true
	// alias
	r.types[ue] = v
	for _, u := range r.Alias.Ptr(elem) {
		r.types[u] = v
	}
	return
}

func (r *Map[V]) MustGetType(elem reflect.Type) (v V) {
	v, _ = r.GetType(elem)
	return
}

func (r *Map[V]) Ptr(in uintptr) (v V, ok bool) {
	v, ok = r.types[in]
	return
}

func (r *Map[V]) Get(in any) V {
	if v, ok := r.types[Ptr(in)]; ok {
		return v
	}
	if v, ok := r.GetType(reflect.TypeOf(in)); ok {
		return v
	}
	panic(ErrValue{in})
}

func (r *Map[V]) Gets(in ...any) []V {
	v := make([]V, 0, len(in))
	for _, i := range in {
		v = append(v, r.Get(i))
	}
	return v
}

func (r *Map[V]) Init(in ...any) *Map[V] {
	r.Gets(in...)
	return r
}

func NewMap[P Parser[V], V any](p P, alias ...Alias) *Map[V] {
	m := &Map[V]{
		Parser: p,
		types:  make(map[uintptr]V),
	}
	if len(alias) == 1 {
		m.Alias = alias[0]
	} else if len(alias) > 1 {
		m.Alias = Aliases(alias)
	} else {
		m.Alias = PTRALIAS
	}
	return m
}
