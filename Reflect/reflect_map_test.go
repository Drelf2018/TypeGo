package Reflect_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/Drelf2018/TypeGo/Reflect"
)

type Tag struct {
	Reflect.Field
	Fields []*Tag
}

func (Tag) Insert(field Reflect.Field, ok bool, fields []*Tag, array *[]*Tag) {
	if !ok {
		return
	}
	*array = slices.Insert(*array, 0, &Tag{
		Field:  field,
		Fields: fields,
	})
}

func TestMap(t *testing.T) {
	tag := Reflect.NewMap[Reflect.TagParser[Tag]]("ref")
	v := tag.Get(Struct1{})
	for idx, val := range v {
		fmt.Printf("#%d: %v\n", idx, val)
	}
	v = tag.Get(&Struct1{})
	for idx, val := range v {
		fmt.Printf("#%d: %v\n", idx, val)
	}
}
