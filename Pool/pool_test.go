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
