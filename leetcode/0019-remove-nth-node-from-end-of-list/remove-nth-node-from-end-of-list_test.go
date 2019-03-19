package prob0019

import (
	"github.com/allbuleyu/algorithms/kit"
	"testing"

	"github.com/stretchr/testify/assert"
	"fmt"
)

func Test_prb0002(t *testing.T)  {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input []int
		input2 int
		ans []int
	}{
		{[]int{1,2,3,4,5}, 2, []int{1,2,3,5}},
		{[]int{1,2,3,4,5}, 5, []int{2,3,4,5}},
		{[]int{1,2,3,4,5}, 1, []int{1,2,3,4}},
		{[]int{1,2}, 1, []int{1}},
		{[]int{1}, 1, []int{}},


	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		actual := kit.TransferNodes(removeNthFromEnd(kit.CreateNodes(tc.input), tc.input2))
		ast.Equal(tc.ans, actual, "输入:%v", tc)
	}
}