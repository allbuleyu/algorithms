package prob0010

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isMatch(t *testing.T) {
	ast := assert.New(t)

	// test case
	tcs := []struct {
		input  string
		input2 string
		ans    bool
	}{
		{"aa", "a", false},
		{"", "a*", true},
		{"aa", "a*", true},
		{"ab", ".*", true},
		{"aaa", "a*b", false},
		{"aaa", "a*", true},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, isMatch(tc.input, tc.input2), "输入:%v", tc)
	}
}
