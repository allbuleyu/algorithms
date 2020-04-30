package _530_week_03_3304

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)


func Test_search(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		input []int
		target int
		ans int
	}{
		// TODO: Add test cases.
		{[]int{4,5,6,7,0,1,2}, 0, 4},
		{[]int{2,4,5,6,7,0,1}, 7, 4},

	}

	for _, tc := range tests {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, search(tc.input, tc.target), "输入:%v", tc)
	}
}

func Test_findRotatedIndex(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		input []int
		ans int
	}{
		// TODO: Add test cases.
		{[]int{4,5,6,7,0,1,2}, 4},
		{[]int{2,4,5,6,7,0,1}, 5},
		{[]int{7,0,1,2,4,5,6}, 1},
		{[]int{2,4,5,6,7}, -1},
		{[]int{2,1}, 1},



	}

	for _, tc := range tests {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findRotatedIndex(tc.input, 0, len(tc.input)), "输入:%v", tc)
	}
}
