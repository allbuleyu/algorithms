package prob0459

import (
	"testing"

	"fmt"
	"github.com/stretchr/testify/assert"
)

func Test_prb0014(t *testing.T)  {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input string
		ans bool
	}{
		{"abac", false},
		{"abab", true},
		{"aba", false},
		{"abcabcabcabc", true},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, repeatedSubstringPattern(tc.input), "输入:%v", tc)
	}
}