package Reflect

import (
	"errors"
	"reflect"
	"unsafe"

	"github.com/Drelf2018/TypeGo/Chan"
)

var (
	ErrValue = errors.New("you should pass in a struct or a pointer to a struct")
)

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

func Type(typ reflect.Type) uintptr {
	return Ptr(reflect.Zero(typ).Interface())
}

func Field(field reflect.StructField) uintptr {
	return Type(field.Type)
}

func fields(typ reflect.Type) Chan.Chan[reflect.StructField] {
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

type Value[V any] struct {
	Ptr uintptr
	Val V
}

type Values[V any] []Value[V]

func (vs Values[V]) Values() []V {
	result := make([]V, 0, len(vs))
	for _, v := range vs {
		result = append(result, v.Val)
	}
	return result
}

type Reflect[V any] struct {
	types map[uintptr]Values[V]
	Parse func(self *Reflect[V], field reflect.StructField) V
}

func (r *Reflect[V]) Clear() {
	r.types = make(map[uintptr]Values[V])
}

func (r *Reflect[V]) Ptr(in uintptr) Values[V] {
	return r.types[in]
}

func (r *Reflect[V]) get(typ reflect.Type) (v Values[V], err error) {
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return nil, ErrValue
	}
	v = make(Values[V], 0, typ.NumField())
	fields(typ).Do(func(field reflect.StructField) {
		ptr := Field(field)
		if fv, err := r.get(field.Type); err == nil {
			r.types[ptr] = fv
		}
		v = append(v, Value[V]{ptr, r.Parse(r, field)})
	})
	return
}

func (r *Reflect[V]) Get(in any) (v Values[V]) {
	ptr := Ptr(in)
	if val, ok := r.types[ptr]; ok {
		return val
	}
	v, err := r.get(reflect.TypeOf(in))
	if err != nil {
		panic(err)
	}
	r.types[ptr] = v
	return
}

func New[V any](parse func(self *Reflect[V], field reflect.StructField) V) *Reflect[V] {
	return &Reflect[V]{
		types: make(map[uintptr]Values[V]),
		Parse: parse,
	}
}

func NewTag(tag string) *Reflect[string] {
	return New(func(self *Reflect[string], field reflect.StructField) string {
		return field.Tag.Get(tag)
	})
}

func NewTagWithName(tag string) *Reflect[[2]string] {
	return New(func(self *Reflect[[2]string], field reflect.StructField) [2]string {
		return [2]string{field.Name, field.Tag.Get(tag)}
	})
}

type Tag struct {
	Tag    string
	Fields []Tag
}

func NewTagStruct(tag string) *Reflect[Tag] {
	return New(func(self *Reflect[Tag], field reflect.StructField) Tag {
		return Tag{Tag: field.Tag.Get(tag), Fields: self.Ptr(Field(field)).Values()}
	})
}
