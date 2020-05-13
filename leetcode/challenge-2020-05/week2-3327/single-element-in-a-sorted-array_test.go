package week2_3327

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_singleNonDuplicate(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		input []int
		ans int
	}{
		// TODO: Add test cases.
		{[]int{1,1,2,3,3,4,4,8,8}, 2},
		{[]int{3,3,7,7,10,11,11}, 10},
		{[]int{3,3,7,7,10,10,11}, 11},
		{[]int{1,3,3,7,7,10,10}, 1},
		{[]int{1}, 1},
	}

	for _, tc := range tests {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, singleNonDuplicate(tc.input), "输入:%v", tc)
	}
}