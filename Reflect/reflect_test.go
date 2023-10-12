package Reflect_test

import (
	"fmt"
	"testing"

	"github.com/Drelf2018/TypeGo/Reflect"
)

type Struct1 struct {
	Struct2 struct {
		d1 string `ref:"1"`
		d2 int64  `ref:"14"`
	} `ref:"514"`

	Struct3 struct {
		d3 bool    `ref:"19"`
		d4 float64 `ref:"19"`
	} `ref:"810"`

	D5 *Struct1 `ref:"Struct1"`

	Struct4 struct {
		d7 *Struct1 `ref:"Struct1"`
	} `ref:"Struct4"`
}

func TestTag(t *testing.T) {
	tag := Reflect.NewTagStruct("ref")
	v := tag.Get(Struct1{})
	for idx, val := range v {
		fmt.Printf("#%d: {Ptr: %v, Val: %v}\n  val.Ptr: %v\n", idx, val.Ptr, val.Val, tag.Ptr(val.Ptr))
	}
}
