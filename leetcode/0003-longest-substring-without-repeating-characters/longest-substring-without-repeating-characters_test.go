package prob0003

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {

	tcs := []struct{
		input string
		ans int
	}{
		{
			"abcabcbb",
			3,
		},
		{
			"pwwkew",
			3,
		},
		{
			"bbbbbb",
			1,
		},
	}


	ast := assert.New(t)
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, lengthOfLongestSubstring(tc.input), "输入:%v", tc)
	}
}