package prob0039

import (
	"testing"

	"fmt"
	"github.com/stretchr/testify/assert"
)

func Test_prb0039(t *testing.T)  {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		target int
		input []int
		ans [][]int
	}{
		{
			7,
			[]int{2,3,6,7},
			[][]int{
				[]int{2,2,3},
				[]int{7},
			},
		},
		{
			8,
			[]int{2,3,5},
			[][]int{
				[]int{2,2,2,2},
				[]int{2,3,3},
				[]int{3,5},
			},
		},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, combinationSum(tc.input, tc.target), "输入:%v", tc)
	}
}