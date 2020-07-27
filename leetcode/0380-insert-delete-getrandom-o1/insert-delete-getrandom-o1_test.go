package prob0380

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	r := Constructor()
	r.Insert(0)
	r.Insert(1)
	r.Remove(0)
	r.Insert(2)
	r.Remove(1)

	v := r.GetRandom()
	fmt.Println(v)
}

func Test2(t *testing.T) {
	r := Constructor()
	r.Insert(1)
	r.Remove(2)
	r.Insert(2)
	v := r.GetRandom()
	fmt.Println(v)
	r.Remove(1)

	v = r.GetRandom()
	fmt.Println(v)
}