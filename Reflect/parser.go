package Reflect

import (
	"reflect"
)

type Parser[V any] interface {
	Parse(*Map[V], reflect.Type) V
}

type Field struct {
	Index int
	Name  string
	Tag   string
}

type TagParser[T interface {
	Insert(field Field, ok bool, fields []*T, array *[]*T)
}] string

func (tag TagParser[T]) Parse(ref *Map[[]*T], elem reflect.Type) []*T {
	var zero T
	v := make([]*T, 0, elem.NumField())
	for index, field := range FieldOf(elem) {
		t, ok := field.Tag.Lookup(string(tag))
		zero.Insert(Field{index, field.Name, t}, ok, ref.MustGetType(field.Type), &v)
	}
	return v
}

type ForFields[V any] func(*Map[[]V], reflect.StructField, reflect.Type) V

func (f ForFields[V]) Parse(ref *Map[[]V], elem reflect.Type) []V {
	v := make([]V, 0, elem.NumField())
	for _, field := range FieldOf(elem) {
		v = append(v, f(ref, field, elem))
	}
	return v
}
