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
	D6 string

	Struct4 struct {
		d7 *Struct1 `ref:"Struct7"`
	} `ref:"Struct4"`
}

func TestTag(t *testing.T) {
	tag := Reflect.NewTagStruct("ref")
	v := tag.Get(&Struct1{})
	for idx, val := range v {
		fmt.Printf("#%d: %v\n", idx, val)
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
#0: Tag(514, [Tag(1), Tag(14)])
#1: Tag(810, [Tag(19), Tag(19)])
#2: Tag(Struct1)
#3: Tag()
#4: Tag(Struct4, [Tag(Struct7)])
--- PASS: TestTag (0.00s)
PASS
ok      github.com/Drelf2018/TypeGo/Reflect     0.031s
```