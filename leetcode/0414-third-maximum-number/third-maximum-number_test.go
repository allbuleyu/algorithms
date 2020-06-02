package prob0414

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ffff(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		input []int
		ans int
	}{
		// TODO: Add test cases.
		{[]int{1,2,2,5,3,5}, 2},
		{[]int{1,2,3,4,5}, 3},
		{[]int{1,2}, 2},
		{[]int{1,2,3}, 1},
		{[]int{1,2,-2147483648}, -2147483648},
		{[]int{1,2,3}, 1},

	}

	for _, tc := range tests {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, ffff(tc.input), "输入:%v", tc)
	}
}