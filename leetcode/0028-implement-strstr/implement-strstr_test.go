package prob0028

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
		needle string
		ans int
	}{
		{"hello", "ll", 2},
		{"aaaaa", "bba", -1},
		{"aaa", "aa", 0},
		{"aba", "ba", 1},
		{"", "", 0},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, strStr(tc.input, tc.needle), "输入:%v", tc)
	}
}