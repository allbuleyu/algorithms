package challenge

import (
	"fmt"
	"testing"
)

func Test_Stack(t *testing.T) {
	s := Constructor()
	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	x := s.GetMin()
	s.Pop()
	x = s.Top()
	x = s.GetMin()

	fmt.Println(x)
}