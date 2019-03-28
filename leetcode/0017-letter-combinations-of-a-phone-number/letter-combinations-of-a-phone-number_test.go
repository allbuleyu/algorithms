package prob0017

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_letterCombinations(t *testing.T) {
	ast := assert.New(t)

	tcs := []struct{
		input string
		ans []string
	}{
		//{"234", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"123", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"213", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"1", []string{}},
		{"", []string{}},

	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, letterCombinations(tc.input), "输入:%v", tc)
	}
}