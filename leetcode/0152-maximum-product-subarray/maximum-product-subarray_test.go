package prob0152

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
			[]int{2,3,-2,4},
			6,
		},
		{
			[]int{2,3,-2,-1},
			12,
		},
		{
			[]int{-2,0,-1},
			0,
		},
		{
			[]int{3,-1,4},
			4,
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
			2,
		},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, maxProduct(tc.input), "输入:%v", tc)
	}
}