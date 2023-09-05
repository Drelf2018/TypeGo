package test

import "fmt"

type Student struct {
	Name string
	ID   int64
}

func (s Student) String() string {
	if s.Name == "" {
		return "I don't have a NAME!"
	}
	return fmt.Sprintf("I am %v and my ID is %v.", s.Name, s.ID)
}

func (s *Student) Set(...any) {}

func (s *Student) Reset() {
	s.Name = "张三"
	s.ID = 19260817
}

func (s Student) Introduce() {
	fmt.Println(s.String())
}
