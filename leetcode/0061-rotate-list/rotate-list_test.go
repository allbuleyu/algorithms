package prob0061

import (
	"fmt"
	"github.com/allbuleyu/algorithms/kit"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_rotateRight(t *testing.T)  {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input []int
		input2 int
		ans []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{4,5,1,2,3}},
		{[]int{1, 2, 3, 4, 5}, 5, []int{1,2,3,4,5}},
		{[]int{1, 2, 3, 4, 5}, 15, []int{1,2,3,4,5}},
		{[]int{0,1,2}, 4, []int{2,0,1}},
		{[]int{1}, 4, []int{1}},
		{[]int{}, 4, []int{}},
		{[]int{1}, 0, []int{1}},

	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		input := kit.CreateNodes(tc.input)

		ast.Equal(tc.ans, kit.TransferNodes(rotateRight(input, tc.input2)), "输入:%v", tc)
	}
}