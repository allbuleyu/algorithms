package challenge

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		input []int
		ans []int
	}{
		// TODO: Add test cases.
		{[]int{0,1,0,3,12}, []int{1,3,12,0,0}},

	}

	for _, tc := range tests {
		fmt.Printf("~~%v~~\n", tc)
		moveZeroes(tc.input)
		ast.Equal(tc.ans, tc.input, "输入:%v", tc)
	}
}