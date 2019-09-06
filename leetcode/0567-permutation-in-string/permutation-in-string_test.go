package prob0567

import (
	"testing"

	"fmt"
	"github.com/stretchr/testify/assert"
)

func Test_checkInclusion(t *testing.T)  {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		s1,s2 string
		ans bool
	}{
		{"ab", "eidbaooo", true},
		{"ab", "eidboaoo", false},
		{"abcdxabcde", "abcdeabcdx", true},

	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, checkInclusion(tc.s1, tc.s2), "输入:%v", tc)
	}
}
