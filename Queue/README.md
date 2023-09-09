# Queue

带类型的队列

### 使用

```go
package Queue_test

import (
	"fmt"
	"testing"

	"github.com/Drelf2018/TypeGo/Queue"
)

func countPrime(n int) (count int) {
	composite := make([]bool, n+1)
	q := Queue.New(2)
	for i := range q.Chan() {
		count++
		for j := 2; i*j <= n; j++ {
			composite[i*j] = true
		}
		for j := i + 1; j <= n; j++ {
			if !composite[j] {
				q.Push(j)
				break
			}
		}
		q.Next()
	}
	return
}

func TestQueue(t *testing.T) {
	fmt.Printf("countPrime(1000): %v\n", countPrime(1000))
	fmt.Printf("countPrime(10000): %v\n", countPrime(10000))
	fmt.Printf("countPrime(100000): %v\n", countPrime(100000))
	fmt.Printf("countPrime(1000000): %v\n", countPrime(1000000))
	fmt.Printf("countPrime(10000000): %v\n", countPrime(10000000))
	fmt.Printf("countPrime(100000000): %v\n", countPrime(100000000))
	fmt.Printf("countPrime(1000000000): %v\n", countPrime(1000000000))
}
```

#### 命令

```go
go test github.com/Drelf2018/TypeGo/Queue -v
```

#### 控制台

```
=== RUN   TestQueue
countPrime(1000): 168
countPrime(10000): 1229
countPrime(100000): 9592
countPrime(1000000): 78498
countPrime(10000000): 664579
countPrime(100000000): 5761455
countPrime(1000000000): 50847534
--- PASS: TestQueue (67.16s)
PASS
ok      github.com/Drelf2018/TypeGo/Queue       67.298s
```