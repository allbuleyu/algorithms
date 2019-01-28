package prob0053

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_prob0053(t *testing.T)  {
	ast := assert.New(t)

	tcs := []struct {
		input []int
		ans int
	}{
		{
			[]int{1,-3,4, -2,-1, 6},
			7,
		},
		{
			[]int{-2,11,-4,13,-5,2},
			20,
		},
		{
			[]int{-2,1,-3,4,-1,2,1,-5,4},
			6,
		},
		{
			[]int{1},
			1,
		},
		{
			[]int{2},
			2,
		},
		// 数组为负的情况没想到
		{
			[]int{-2, -1},
			-1,
		},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, maxSubArray(tc.input), "输入:%v", tc)
	}
}