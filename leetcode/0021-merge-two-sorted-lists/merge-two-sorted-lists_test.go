package prob0021

import (
	"algorithms/kit"
	"testing"

	"fmt"
	"github.com/stretchr/testify/assert"
)

func Test_prb00201(t *testing.T) {
	ast := assert.New(t)

	// test case
	tcs := []struct {
		input  []int
		input2 []int
		ans    []int
	}{
		{[]int{1, 2, 3}, []int{0, 2, 3, 4, 6, 6}, []int{0, 1, 2, 2, 3, 3, 4, 6, 6}},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		actual := kit.TransferNodes(mergeTwoLists(kit.CreateNodes(tc.input), kit.CreateNodes(tc.input2)))
		ast.Equal(tc.ans, actual, "输入:%v", tc)
	}
}

type Abc struct {
	m map[string]string
}

func Test_One(t *testing.T) {
	abc := &Abc{}
	_, ok := abc.m["abc"]
	if !ok {
		t.Log("not panic")
	}

	var m map[string]string
	_, ok = m["abc"]
	if !ok {
		t.Log("not panic")
	}
}
