# Pool

带类型的变量池

### 使用

```go
package Pool_test

import (
	"testing"

	"github.com/Drelf2018/TypeGo/Pool"
	"github.com/Drelf2018/TypeGo/test"
)

func TestPool(t *testing.T) {
	pool := Pool.New(&test.Student{})

	s1 := pool.Get("张三")
	t.Logf("s1: %v\n", s1)

	pool.Put(s1)
	t.Logf("s1: %v\n", s1)

	s2 := pool.Get("李四")
	t.Logf("s2: %v\n", s2)

	s3 := pool.Get("王五")
	t.Logf("s3: %v\n", s3)
}
```

#### 命令

```go
go test github.com/Drelf2018/TypeGo/Pool -v
```

#### 控制台

```
=== RUN   TestPool
    pool_test.go:14: s1: I am 张三 and my ID is 1.
    pool_test.go:17: s1: I am undefined and my ID is 0.
    pool_test.go:20: s2: I am 李四 and my ID is 0.
    pool_test.go:23: s3: I am 王五 and my ID is 1.
--- PASS: TestPool (0.00s)
PASS
ok      github.com/Drelf2018/TypeGo/Pool        0.028s
```