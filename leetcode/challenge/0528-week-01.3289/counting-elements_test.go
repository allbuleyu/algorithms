package _528_week_01_3289

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
		{[]int{1,2,3}, 2},
		{[]int{1,1,3,3,5,5,7,7}, 0},
		{[]int{1,1,2,2}, 2},

	}

	for _, tc := range tests {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, countElements(tc.input), "输入:%v", tc)
	}
}