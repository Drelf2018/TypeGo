# Reflect

更好的反射

### 使用

```go
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
```

#### 命令

```go
go test github.com/Drelf2018/TypeGo/Reflect -v
```

#### 控制台

```
=== RUN   TestTag
#0: {Ptr: 9254752, Val: {514 [{1 []} {14 []}]}}
  val.Ptr: [{9198976 {1 []}} {9199552 {14 []}}]
#1: {Ptr: 9254880, Val: {810 [{19 []} {19 []}]}}
  val.Ptr: [{9199040 {19 []}} {9199296 {19 []}}]
#2: {Ptr: 9301280, Val: {Struct1 [{514 [{1 []} {14 []}]} {810 [{19 []} {19 []}]}]}}
  val.Ptr: [{9254752 {514 [{1 []} {14 []}]}} {9254880 {810 [{19 []} {19 []}]}} {9301280 {Struct1 [{514 [{1 []} {14 []}]} {810 [{19 []} {19 []}]}]}} {9242176 {Struct4 [{Struct1 [{514 [{1 []} {14 []}]} {810 [{19 []} {19 []}]} {Struct1 [{514 [{1 []} {14 []}]} {810 [{19 []} {19 []}]}]}]}]}}]
#3: {Ptr: 9242176, Val: {Struct4 [{Struct1 [{514 [{1 []} {14 []}]} {810 [{19 []} {19 []}]} {Struct1 [{514 [{1 []} {14 []}]} {810 [{19 []} {19 []}]}]}]}]}}
  val.Ptr: [{9301280 {Struct1 [{514 [{1 []} {14 []}]} {810 [{19 []} {19 []}]} {Struct1 [{514 [{1 []} {14 []}]} {810 [{19 []} {19 []}]}]}]}}]
--- PASS: TestTag (0.00s)
PASS
ok      github.com/Drelf2018/TypeGo/Reflect     0.027s
```