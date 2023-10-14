package Reflect

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"unsafe"

	"github.com/Drelf2018/TypeGo/Chan"
)

var ErrValue = errors.New("you should pass in a struct or a pointer to a struct")

// 反射加速 参考: https://www.cnblogs.com/cheyunhua/p/16642488.html
type eface struct {
	Type  unsafe.Pointer
	Value unsafe.Pointer
}

func (e *eface) Ptr() uintptr {
	return uintptr(e.Type)
}

func Eface(in any) *eface {
	return (*eface)(unsafe.Pointer(&in))
}

func Ptr(in any) uintptr {
	return Eface(in).Ptr()
}

func Addr(typ reflect.Type) uintptr {
	return Ptr(reflect.New(typ).Interface())
}

func Type(typ reflect.Type) uintptr {
	return Ptr(reflect.Zero(typ).Interface())
}

func fields(typ reflect.Type) Chan.Chan[reflect.StructField] {
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	return Chan.Auto(func(c chan reflect.StructField) {
		for i, l := 0, typ.NumField(); i < l; i++ {
			c <- typ.Field(i)
		}
	})
}

func Fields(v any) Chan.Chan[reflect.StructField] {
	typ := reflect.TypeOf(v)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		panic(ErrValue)
	}
	return fields(typ)
}

type Reflect[V any] struct {
	types map[uintptr][]V
	Parse func(self *Reflect[V], field reflect.StructField) V
}

func (r *Reflect[V]) Clear() {
	r.types = make(map[uintptr][]V)
}

func (r *Reflect[V]) ptr(in any, v *[]V) (ok bool) {
	return r.Ptr(Ptr(in), v)
}

func (r *Reflect[V]) Ptr(in uintptr, v *[]V) (ok bool) {
	if v == nil {
		_, ok = r.types[in]
		return
	}
	*v, ok = r.types[in]
	return
}

func (r *Reflect[V]) GetType(elem reflect.Type, v *[]V) bool {
	if elem.Kind() == reflect.Ptr {
		elem = elem.Elem()
	}
	if elem.Kind() != reflect.Struct {
		return false
	}
	ue := Type(elem)
	if r.Ptr(ue, v) {
		return true
	}
	r.types[ue] = make([]V, 0)
	values := make([]V, 0, elem.NumField())
	fields(elem).Do(func(field reflect.StructField) {
		values = append(values, r.Parse(r, field))
	})
	r.types[ue] = values
	r.types[Addr(elem)] = values
	return r.Ptr(ue, v)
}

func (r *Reflect[V]) Get(in any) (v []V) {
	if r.ptr(in, &v) || r.GetType(reflect.TypeOf(in), &v) {
		return
	}
	panic(ErrValue)
}

func New[V any](parse func(self *Reflect[V], field reflect.StructField) V) *Reflect[V] {
	return &Reflect[V]{
		types: make(map[uintptr][]V),
		Parse: parse,
	}
}

func NewTag(tag string) *Reflect[string] {
	return New(func(self *Reflect[string], field reflect.StructField) string {
		return field.Tag.Get(tag)
	})
}

type Tag struct {
	Tag    string
	Fields []Tag
}

func (t Tag) String() string {
	return fmt.Sprintf("Tag(%v%v)", t.Tag, Tags(t.Fields))
}

type Tags []Tag

func (t Tags) String() string {
	l := len(t)
	if l == 0 {
		return ""
	}
	buf := bytes.NewBufferString(", [")
	for i, f := range t {
		buf.WriteString(f.String())
		if i != l-1 {
			buf.WriteString(", ")
		}
	}
	buf.WriteString("]")
	return buf.String()
}

func NewTagStruct(tag string) *Reflect[Tag] {
	return New(func(self *Reflect[Tag], field reflect.StructField) (t Tag) {
		t.Tag = field.Tag.Get(tag)
		self.GetType(field.Type, &t.Fields)
		return
	})
}
