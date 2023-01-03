package prob0206

import (
	"algorithms/kit"
	"testing"

	"fmt"
	"github.com/stretchr/testify/assert"
)

func Test_prb00206(t *testing.T) {
	ast := assert.New(t)

	// test case
	tcs := []struct {
		input []int

		ans []int
	}{
		{[]int{1, 2}, []int{2, 1}},
		{[]int{1, 2, 3}, []int{3, 2, 1}},
		{[]int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		actual := kit.TransferNodes(reverseList(kit.CreateNodes(tc.input)))
		ast.Equal(tc.ans, actual, "输入:%v", tc)
	}
}

type ttt struct {
	v    int
	next *ttt
}

func Abc(t *ttt) *ttt {
	x := &ttt{v: 111}
	t = x

	return t
}

func Abcd(t *ttt) {
	x := &ttt{v: 111}
	*t = *x
}

func Test_Pointer(t *testing.T) {
	x := &ttt{v: 100}
	x = Abc(x)

	if x.v != 111 {
		t.Error("wrong way to use pointer")
	}
}
