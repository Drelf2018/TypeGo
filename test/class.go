package test

type Class struct {
	Students []Student
}

func (c *Class) Join(s Student) {
	c.Students = append(c.Students, s)
}

func (c Class) Call(ch chan Student) {
	for _, s := range c.Students {
		ch <- s
	}
	close(ch)
}
