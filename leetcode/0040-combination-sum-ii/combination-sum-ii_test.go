package prob0040

import (
"testing"

"fmt"
"github.com/stretchr/testify/assert"
)

func Test_prb0040(t *testing.T)  {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		target int
		input []int
		ans [][]int
	}{
		{
			5,
			[]int{2,5,2,1,2},
			[][]int{
				[]int{1,2,2},
				[]int{5},
			},
		},
		//{
		//	8,
		//	[]int{10,1,2,7,6,1,5},
		//	[][]int{
		//		[]int{1,7},
		//		[]int{1,2,5},
		//		[]int{2,6},
		//		[]int{1, 1, 6},
		//	},
		//},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, combinationSum2(tc.input, tc.target), "输入:%v", tc)
	}
}