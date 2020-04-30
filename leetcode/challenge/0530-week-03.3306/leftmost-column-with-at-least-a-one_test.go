package _530_week_03_3306

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findFirstIndex(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		input []int
		ans int
	}{
		// TODO: Add test cases.
		{[]int{0,1,1,1,1,1,1}, 1},
		{[]int{0,0,0,0,0,0,0,0,0,1,1}, 9},
		{[]int{0,0,1}, 2},
		{[]int{0,0,0}, -1},

	}

	for _, tc := range tests {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findFirstIndex(tc.input), "输入:%v", tc)
	}
}