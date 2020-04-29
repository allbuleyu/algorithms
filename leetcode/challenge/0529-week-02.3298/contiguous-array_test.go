package _529_week_02_3298

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_subarraySum(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		input []int
		ans int
	}{
		// TODO: Add test cases.
		{[]int{0,1,1,0,1,1,1,0}, 4},
		{[]int{0,1}, 2},
		{[]int{0, 1, 0}, 2},

	}

	for _, tc := range tests {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findMaxLength(tc.input), "输入:%v", tc)
	}
}
