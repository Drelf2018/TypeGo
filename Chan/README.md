# Chan

带类型的单向通道

### 使用

```go
package Chan_test

import (
	"fmt"
	"testing"

	"github.com/Drelf2018/TypeGo/Chan"
	"github.com/Drelf2018/TypeGo/test"
)

func TestChan(t *testing.T) {
	class := test.Class{}
	class.Join(test.Student{Name: "张三", ID: 1})
	class.Join(test.Student{Name: "李四", ID: 2})
	class.Join(test.Student{Name: "王五", ID: 3})

	for s := range Chan.New(class.Call) {
		s.Introduce()
	}
}

func TestSlice(t *testing.T) {
	for i := range Chan.Slice([]int{1, 1, 4, 5, 1, 4}) {
		print(i)
	}
	println()
}

func TestMap(t *testing.T) {
	m := map[string]int{
		"Alice": 1,
		"Bob":   2,
		"Carol": 3,
	}
	for k := range Chan.Keys(m) {
		print(k, " ")
	}
	println()
	for v := range Chan.Values(m) {
		print(v, " ")
	}
	println()
}

// 参考: https://www.runoob.com/python/python-func-range.html
func TestRange(t *testing.T) {
	// >>>range(10)        # 从 0 开始到 9
	// [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
	fmt.Printf("Chan.Range(10).List(): %v\n", Chan.Range(10).List())

	// >>> range(1, 11)     # 从 1 开始到 10
	// [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
	fmt.Printf("Chan.Range(1, 11).List(): %v\n", Chan.Range(1, 11).List())

	// >>> range(0, 30, 5)  # 步长为 5
	// [0, 5, 10, 15, 20, 25]
	fmt.Printf("Chan.Range(0, 30, 5).List(): %v\n", Chan.Range(0, 30, 5).List())

	// >>> range(0, 10, 3)  # 步长为 3
	// [0, 3, 6, 9]
	fmt.Printf("Chan.Range(0, 10, 3).List(): %v\n", Chan.Range(0, 10, 3).List())

	// >>> range(0, -10, -1) # 负数
	// [0, -1, -2, -3, -4, -5, -6, -7, -8, -9]
	fmt.Printf("Chan.Range(0, -10, -1).List(): %v\n", Chan.Range(0, -10, -1).List())

	// >>> range(0)
	// []
	fmt.Printf("Chan.Range(0).List(): %v\n", Chan.Range(0).List())

	// >>> range(1, 0)
	// []
	fmt.Printf("Chan.Range(0).List(): %v\n", Chan.Range(1, 0).List())
}
```

#### 命令

```go
go test github.com/Drelf2018/TypeGo/Chan -v
```

#### 控制台

```
=== RUN   TestChan
I am 张三 and my ID is 1.
I am 李四 and my ID is 2.
I am 王五 and my ID is 3.
--- PASS: TestChan (0.00s)
=== RUN   TestSlice
114514
--- PASS: TestSlice (0.00s)
=== RUN   TestMap
Alice Bob Carol
1 2 3
--- PASS: TestMap (0.00s)
=== RUN   TestRange
Chan.Range(10).List(): [0 1 2 3 4 5 6 7 8 9]
Chan.Range(0, 30, 5).List(): [0 5 10 15 20 25]
Chan.Range(0, 10, 3).List(): [0 3 6 9]
Chan.Range(0, -10, -1).List(): [0 -1 -2 -3 -4 -5 -6 -7 -8 -9]
Chan.Range(0).List(): []
Chan.Range(0).List(): []
--- PASS: TestRange (0.00s)
PASS
ok      github.com/Drelf2018/TypeGo/Chan        0.052s
```