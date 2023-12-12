package Reflect

import (
	"bytes"
	"fmt"
	"reflect"
	"unsafe"

	"github.com/Drelf2018/TypeGo/Chan"
)

type ErrValue struct {
	in any
}

func (e ErrValue) Error() string {
	return fmt.Sprintf("you need to pass in (a pointer to) a struct instead of %#v", e.in)
}

// 反射加速
//
// 参考: https://www.cnblogs.com/cheyunhua/p/16642488.html
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

func Fields(elem reflect.Type) Chan.Chan[reflect.StructField] {
	if elem.Kind() == reflect.Ptr {
		elem = elem.Elem()
	}
	if elem.Kind() != reflect.Struct {
		panic(ErrValue{elem})
	}
	return Chan.Auto(func(c chan reflect.StructField) {
		for i, l := 0, elem.NumField(); i < l; i++ {
			c <- elem.Field(i)
		}
	})
}

func FieldOf(elem reflect.Type) []reflect.StructField {
	if elem.Kind() == reflect.Ptr {
		elem = elem.Elem()
	}
	if elem.Kind() != reflect.Struct {
		panic(ErrValue{elem})
	}
	l := elem.NumField()
	r := make([]reflect.StructField, 0, l)
	for i := 0; i < l; i++ {
		r = append(r, elem.Field(i))
	}
	return r
}

func MethodOf(value reflect.Value) map[reflect.Method]reflect.Value {
	if value.Kind() != reflect.Struct && (value.Kind() != reflect.Ptr || value.Elem().Kind() != reflect.Struct) {
		panic(ErrValue{value})
	}
	elem := value.Type()
	r := make(map[reflect.Method]reflect.Value)
	for i := elem.NumMethod() - 1; i >= 0; i-- {
		r[elem.Method(i)] = value.Method(i)
	}
	return r
}

func MethodFuncOf(value reflect.Value) map[string]any {
	if value.Kind() != reflect.Struct && (value.Kind() != reflect.Ptr || value.Elem().Kind() != reflect.Struct) {
		panic(ErrValue{value})
	}
	elem := value.Type()
	r := make(map[string]any)
	for i := elem.NumMethod() - 1; i >= 0; i-- {
		r[elem.Method(i).Name] = value.Method(i).Interface()
	}
	return r
}

type Reflect[V any] struct {
	types map[uintptr][]V
	Alias func(elem reflect.Type) []uintptr
	Parse func(self *Reflect[V], field reflect.StructField, elem reflect.Type) V
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
	for field := range Fields(elem) {
		values = append(values, r.Parse(r, field, elem))
	}
	r.types[ue] = values
	for _, u := range r.Alias(elem) {
		r.types[u] = values
	}
	return r.Ptr(ue, v)
}

func (r *Reflect[V]) Init(in ...any) {
	for _, i := range in {
		r.GetType(reflect.TypeOf(i), nil)
	}
}

func (r *Reflect[V]) Get(in any) (v []V) {
	if r.ptr(in, &v) || r.GetType(reflect.TypeOf(in), &v) {
		return
	}
	panic(ErrValue{in})
}

func Pointer(elem reflect.Type) []uintptr {
	return []uintptr{
		Addr(elem),
	}
}

func Slice(elem reflect.Type) []uintptr {
	return []uintptr{
		Addr(elem),
		Type(reflect.SliceOf(elem)),
		Addr(reflect.SliceOf(elem)),
	}
}

func SlicePtr(elem reflect.Type) []uintptr {
	return []uintptr{
		Addr(elem),
		Addr(reflect.SliceOf(elem)),
		Addr(reflect.SliceOf(reflect.PtrTo(elem))),
	}
}

type Option[V any] func(*Reflect[V])

func WithAlias[V any](f func(elem reflect.Type) []uintptr) Option[V] {
	return func(r *Reflect[V]) {
		r.Alias = f
	}
}

func WithSlice[V any](in ...any) Option[V] {
	return func(r *Reflect[V]) {
		r.Alias = Slice
		r.Init(in...)
	}
}

func New[V any](parse func(self *Reflect[V], field reflect.StructField, elem reflect.Type) V, options ...Option[V]) (r *Reflect[V]) {
	r = &Reflect[V]{
		types: make(map[uintptr][]V),
		Alias: Pointer,
		Parse: parse,
	}
	for _, op := range options {
		op(r)
	}
	return
}

func NewTag(tag string, options ...Option[string]) *Reflect[string] {
	return New(func(self *Reflect[string], field reflect.StructField, elem reflect.Type) string {
		return field.Tag.Get(tag)
	}, options...)
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

func NewTagStruct(tag string, options ...Option[Tag]) *Reflect[Tag] {
	return New(func(self *Reflect[Tag], field reflect.StructField, elem reflect.Type) (t Tag) {
		t.Tag = field.Tag.Get(tag)
		self.GetType(field.Type, &t.Fields)
		return
	}, options...)
}
