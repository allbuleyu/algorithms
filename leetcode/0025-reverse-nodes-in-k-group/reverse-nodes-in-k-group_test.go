package prob0025

import (
	"github.com/allbuleyu/algorithms/kit"
	"testing"

	"github.com/stretchr/testify/assert"
	"fmt"
)

func Test_swapPairs(t *testing.T)  {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input []int
		input2 int
		ans []int
	}{
		{[]int{1,2,3,4,5}, 2, []int{2,1,4,3,5}},
		{[]int{1,2,3,4,5}, 3, []int{3,2,1,4,5}},


	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		actual := kit.TransferNodes(reverseKGroup(kit.CreateNodes(tc.input), tc.input2))
		ast.Equal(tc.ans, actual, "输入:%v", tc)
	}
}