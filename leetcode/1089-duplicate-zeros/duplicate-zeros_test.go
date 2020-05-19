package prob1089

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_duplicateZeros(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		input []int
		ans []int
	}{
		// TODO: Add test cases.
		//{[]int{1,0,2,3,0,4,5,0}, []int{1,0,0,2,3,0,0,4}},
		//{[]int{1,2,3}, []int{1,2,3}},
		//{[]int{0,0,0,1,2,3,4,5,6}, []int{0,0,0,0,0,0,1,2,3}},
		{[]int{8,4,5,0,0,0,0,7}, []int{8,4,5,0,0,0,0,0}},
	}

	for _, tc := range tests {
		fmt.Printf("~~%v~~\n", tc)
		duplicateZeros(tc.input)
		ast.Equal(tc.ans, tc.input, "输入:%v", tc)
	}
}