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
