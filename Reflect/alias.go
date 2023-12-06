package Reflect

import "reflect"

const (
	NOALIAS       NoAlias       = 1 << iota
	PTRALIAS      PtrAlias      = 1 << iota
	SLICEALIAS    SliceAlias    = 1 << iota
	SLICEPTRALIAS SlicePtrAlias = 1 << iota
)

type Alias interface {
	Ptr(reflect.Type) []uintptr
}

type NoAlias int

func (NoAlias) Ptr(elem reflect.Type) []uintptr {
	return nil
}

type PtrAlias int

func (PtrAlias) Ptr(elem reflect.Type) []uintptr {
	return []uintptr{Addr(elem)}
}

type SliceAlias int

func (SliceAlias) Ptr(elem reflect.Type) []uintptr {
	return []uintptr{
		Addr(elem),
		Type(reflect.SliceOf(elem)),
		Addr(reflect.SliceOf(elem)),
	}
}

type SlicePtrAlias int

func (SlicePtrAlias) Ptr(elem reflect.Type) []uintptr {
	return []uintptr{
		Addr(elem),
		Addr(reflect.SliceOf(elem)),
		Addr(reflect.SliceOf(reflect.PtrTo(elem))),
	}
}

type Aliases []Alias

func (as Aliases) Ptr(elem reflect.Type) (u []uintptr) {
	for _, a := range as {
		u = append(u, a.Ptr(elem)...)
	}
	return u
}
