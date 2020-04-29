package _529_week_01_3299

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_stringShift(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		input string
		shift [][]int
		ans string
	}{
		// TODO: Add test cases.
		{"abc", [][]int{[]int{0,1}, []int{1,2}}, "cab"},
		{"abc", [][]int{[]int{0,1}, []int{1,2}}, "cab"},
		{"abcdefg", [][]int{[]int{1,1}, []int{1,1},[]int{0,2}, []int{1,3}}, "efgabcd"},


	}

	for _, tc := range tests {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, stringShift(tc.input, tc.shift), "输入:%v", tc)
	}
}