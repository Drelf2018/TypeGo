package Queue

import (
	"sync"

	"github.com/Drelf2018/TypeGo/Chan"
)

type Queue[T any] struct {
	items []T
	lock  sync.Mutex
	pop   chan any
}

func New[T any](items ...T) *Queue[T] {
	return &Queue[T]{items: items}
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue[T]) Push(items ...T) {
	q.lock.Lock()
	q.items = append(q.items, items...)
	q.lock.Unlock()
}

func (q *Queue[T]) Pop() (item T, ok bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}
	q.lock.Lock()
	item, q.items = q.items[0], q.items[1:]
	q.lock.Unlock()
	return item, true
}

func (q *Queue[T]) MustPop() T {
	item, _ := q.Pop()
	return item
}

func (q *Queue[T]) Peek() (item T, ok bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}
	return q.items[0], true
}

func (q *Queue[T]) MustPeek() T {
	item, _ := q.Peek()
	return item
}

func (q *Queue[T]) Back() (item T, ok bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}
	return q.items[len(q.items)-1], true
}

func (q *Queue[T]) MustBack() T {
	item, _ := q.Back()
	return item
}

func (q *Queue[T]) Next(items ...T) {
	if len(items) != 0 {
		q.Push(items...)
	}
	q.pop <- struct{}{}
}

func (q *Queue[T]) Chan() Chan.Chan[T] {
	q.pop = make(chan any)
	defer q.Next()
	return Chan.Auto(func(ch chan T) {
		for range q.pop {
			if item, ok := q.Pop(); ok {
				ch <- item
			} else {
				break
			}
		}
	})
}
