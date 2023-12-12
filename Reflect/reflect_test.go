package Reflect_test

import (
	"fmt"
	"reflect"
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
	D6 string

	Struct4 struct {
		d7 *Struct1 `ref:"Struct7"`
	} `ref:"Struct4"`
}

func (*Struct1) Test1(int)   {}
func (Struct1) Test2(string) {}
func (*Struct1) Test3(bool)  {}
func (Struct1) Test4(a, b int) float64 {
	return float64(a + b)
}

func TestMethod(t *testing.T) {
	s := reflect.ValueOf(Struct1{})
	for name, fn := range Reflect.MethodFuncOf(s) {
		switch fn.(type) {
		case func(string):
			fmt.Printf("string: %v\n", name)
		case func(int, int) float64:
			fmt.Printf("float64: %v\n", name)
		}
	}
}

func TestTag(t *testing.T) {
	tag := Reflect.NewTagStruct("ref", Reflect.WithSlice[Reflect.Tag](Struct1{}))
	v := tag.Get(&[]Struct1{})
	for idx, val := range v {
		fmt.Printf("#%d: %v\n", idx, val)
	}
}
