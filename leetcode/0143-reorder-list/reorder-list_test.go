package prob0143

import (
	"fmt"
	"github.com/allbuleyu/algorithms/kit"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reorderList(t *testing.T)  {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input []int
		ans []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1,5,2,4,3}},
		{[]int{1, 2, 3, 4}, []int{1,4,2,3}},
		{[]int{1, 2}, []int{1,2}},
		{[]int{}, []int{}},


	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		input := kit.CreateNodes(tc.input)
		reorderList(input)
		ast.Equal(tc.ans, kit.TransferNodes(input), "输入:%v", tc)
	}
}