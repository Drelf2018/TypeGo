package Chan

// It is just a typed Channel
type Chan[T any] <-chan T

func (c Chan[T]) List() (r []T) {
	for v := range c {
		r = append(r, v)
	}
	return
}

// After sending all your datas, you need to close(ch) manually.
//
// Or you can use Auto() which wiil run close(ch) automatic after f(ch) done.
func New[T any](f func(ch chan T)) Chan[T] {
	ch := make(chan T)
	go f(ch)
	return ch
}

// Run close(ch) automatic after f(ch) done.
func Auto[T any](f func(ch chan T)) Chan[T] {
	return New(func(ch chan T) {
		defer close(ch)
		f(ch)
	})
}

// A Range(start, stop[, step]) function like python range()
func Range[T int | int64 | float64](stop T, options ...T) Chan[T] {
	var start, step, alpha T = 0, 1, 1
	switch len(options) {
	case 2:
		step = options[1]
		fallthrough
	case 1:
		start = stop
		stop = options[0]
	}
	if step < 0 {
		alpha = -1
	}
	return New(func(c chan T) {
		for i := start; (i-stop)*alpha < 0; i += step {
			c <- i
		}
		close(c)
	})
}

func Slice[T any](s []T) Chan[T] {
	return New(func(c chan T) {
		for i, l := 0, len(s); i < l; i++ {
			c <- s[i]
		}
		close(c)
	})
}

func Values[M ~map[K]V, K comparable, V any](m M) Chan[V] {
	return New(func(c chan V) {
		for _, v := range m {
			c <- v
		}
		close(c)
	})
}
