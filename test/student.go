package test

import (
	"fmt"

	"github.com/Drelf2018/TypeGo/Pool"
)

type Student struct {
	Pool.None
	Name string
	ID   int64
}

func (s Student) String() string {
	if s.Name == "" {
		return "I don't have a NAME!"
	}
	return fmt.Sprintf("I am %v and my ID is %v.", s.Name, s.ID)
}

func (s *Student) New() {
	s.ID = 1
}

func (s *Student) Set(x ...any) {
	s.Name = x[0].(string)
}

func (s *Student) Reset() {
	s.Name = "undefined"
	s.ID = 0
}

func (s Student) Introduce() {
	fmt.Println(s.String())
}
