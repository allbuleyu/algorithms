package prob0002

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
		input2 []int
		ans []int
	}{
		{[]int{2,4,3}, []int{5,6,4}, []int{7,0,8}},
		{[]int{9,9}, []int{1}, []int{0,0,1}},
		{[]int{1}, []int{9,9}, []int{0,0,1}},
		{[]int{1,2,3}, []int{2,8,7}, []int{3,0,1,1}},


	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		actual := kit.TransferNodes(addTwoNumbers(kit.CreateNodes(tc.input), kit.CreateNodes(tc.input2)))
		ast.Equal(tc.ans, actual, "输入:%v", tc)
	}
}