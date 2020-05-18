package week3_3332

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findAnagrams(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		input string
		p string
		ans []int
	}{
		// TODO: Add test cases.
		{"cbaebabacd", "abc", []int{0,6}},
		{"abab", "ab", []int{0,1,2}},
	}

	for _, tc := range tests {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, findAnagrams(tc.input, tc.p), "输入:%v", tc)
	}
}