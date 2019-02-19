package prob0746

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"fmt"
)

func Test_prb0001(t *testing.T)  {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input []int
		ans int
	}{
		{
			[]int{10, 15, 20},
			15,
		},
		{
			[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			6,
		},
		{
			[]int{0,0,0,0},
			0,
		},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, minCostClimbingStairs(tc.input), "输入:%v", tc)
	}
}