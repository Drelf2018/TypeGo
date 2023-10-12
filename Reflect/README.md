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

	D5 string `ref:"Reflect"`
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
#0: {Ptr: 16922208, Val: {514 [{1 []} {14 []}]}}
  val.Ptr: [{16866560 {1 []}} {16867136 {14 []}}]
#1: {Ptr: 16922336, Val: {810 [{19 []} {19 []}]}}
  val.Ptr: [{16866624 {19 []}} {16866880 {19 []}}]
#2: {Ptr: 16866560, Val: {Reflect []}}
  val.Ptr: []
--- PASS: TestTag (0.00s)
PASS
ok      github.com/Drelf2018/TypeGo/Reflect     0.027s
```