package prob0001

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"fmt"
)

func Test_prb0001(t *testing.T)  {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		target int
		input, ans []int
	}{
		{
			9,
			[]int{2, 7, 11, 15},
			[]int{0, 1},
		},
		{
			13,
			[]int{2, 7, 11, 15},
			[]int{0, 2},
		},
		{
			18,
			[]int{2, 7, 11, 15},
			[]int{1, 2},
		},
		{
			26,
			[]int{2, 7, 11, 15},
			[]int{2, 3},
		},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, twoSum(tc.input, tc.target), "输入:%v", tc)
	}
}
