package week3_3330

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxSubarraySumCircular(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		input []int
		ans int
	}{
		// TODO: Add test cases.
		{[]int{-1,3,-3,9,-6,8,-5,-5,-6,10}, 20},
		{[]int{0,5,8,-9,9,-7,3,-2}, 16},
		{[]int{-3,3,-1,9}, 11},
		{[]int{1,-2,3,-2}, 3},
		{[]int{1,2,-4,5}, 8},
		{[]int{5,-3,5}, 10},
		{[]int{1}, 1},
		{[]int{-11,-12,-1}, -1},
		{[]int{-11,-12,-1}, -1},
	}

	for _, tc := range tests {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, maxSubarraySumCircular(tc.input), "输入:%v", tc)
	}
}