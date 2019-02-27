package prob0746

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_prob0236(t *testing.T) {
	ast := assert.New(t)

	tcs := []struct {
		input []int
		ans int
	}{
		{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},6,},
		{[]int{10, 15, 20},15,},

	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, minCostClimbingStairs(tc.input), "输入:%v", tc)
	}
}