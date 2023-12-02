# Pool

带类型的变量池

### 使用

```go
package Pool_test

import (
	"errors"
	"testing"

	"github.com/Drelf2018/TypeGo/Pool"
)

type Class struct {
	Students map[int]string
}

func (c *Class) OnNew() {
	c.Students = make(map[int]string)
}

func (c *Class) OnPut() {
	for k := range c.Students {
		delete(c.Students, k)
	}
}

func (c *Class) OnGet(stu ...any) {
	l := len(stu)
	if l&1 == 1 {
		panic(errors.New("odd params"))
	}
	for i := 0; i < l; i += 2 {
		c.Students[stu[i].(int)] = stu[i+1].(string)
	}
}

func TestPool(t *testing.T) {
	pool := Pool.New(new(Class))

	c1 := pool.Get(1, "Alice")
	t.Logf("c1: %v\n", c1)

	pool.Put(c1)
	t.Logf("c1: %v\n", c1)

	c2 := pool.Get(2, "Bob", 3, "Carol")
	t.Logf("c1: %v c2: %v\n", c1, c2)

	c3 := pool.Get(4, "Dave")
	t.Logf("c3: %v\n", c3)

	pool.Put(c2, c3, nil)
	t.Logf("pool.Len(): %v pool.Cap(): %v\n", pool.Len(), pool.Cap())

	_ = pool.Get()
	t.Logf("pool.Len(): %v pool.Cap(): %v\n", pool.Len(), pool.Cap())
}
```

#### 命令

```go
go test github.com/Drelf2018/TypeGo/Pool -v
```

#### 控制台

```
=== RUN   TestPool
    pool_test.go:38: c1: &{map[1:Alice]}
    pool_test.go:41: c1: &{map[]}
    pool_test.go:44: c1: &{map[2:Bob 3:Carol]} c2: &{map[2:Bob 3:Carol]}
    pool_test.go:47: c3: &{map[4:Dave]}
    pool_test.go:50: pool.Len(): 0 pool.Cap(): 2
    pool_test.go:53: pool.Len(): 1 pool.Cap(): 2
--- PASS: TestPool (0.00s)
PASS
ok      github.com/Drelf2018/TypeGo/Pool        0.026s
```