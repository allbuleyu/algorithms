package challenge

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxSubArray(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		input []int
		ans int
	}{
		// TODO: Add test cases.
		{[]int{-2,1,-3,4,-1,2,1,-5,4}, 6},
		{[]int{-2,1,-3,4,-1,2,1,4}, 10},
		{[]int{1}, 1},
		{[]int{-1, -1}, -1},
	}

	for _, tc := range tests {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, maxSubArray(tc.input), "输入:%v", tc)
	}
}