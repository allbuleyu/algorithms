package prob0018

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_fourSum(t *testing.T) {

	tcs := []struct{
		input []int
		target int
		ans [][]int
	}{
		{
			[]int{1, 0, -1, 0, -2, 2},0,
			[][]int{
				[]int{-1,  0, 0, 1},
				[]int{-2, -1, 1, 2},
				[]int{-2,  0, 0, 2},
			},
		},
		{
			[]int{0,0,0,0},0,
			[][]int{
				{0,  0, 0, 0},
			},
		},
		{
			[]int{-3,-1,0,2,4,5},0,
			[][]int{
				{-3,-1,0,4},
			},
		},
	}

	ast := assert.New(t)
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, fourSum(tc.input, tc.target), "输入:%v", tc)
	}
}