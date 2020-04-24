package _531_week_04_3307

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_subarraySum(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		input []int
		k int
		ans int
	}{
		// TODO: Add test cases.
		{[]int{1,1,1}, 2, 2},
		{[]int{0,0,0,0,2}, 2, 5},
		{[]int{0,0,2,0,0}, 2, 9},
	}

	for _, tc := range tests {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, subarraySum(tc.input, tc.k), "输入:%v", tc)
	}
}