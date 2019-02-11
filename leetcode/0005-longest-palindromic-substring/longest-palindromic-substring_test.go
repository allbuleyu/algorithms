package prob0005

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_prob0005(t *testing.T)  {
	ast := assert.New(t)

	tcs := []struct {
		input string
		ans string
	}{
		{
			"babad",
			"bab",
		},
		{
			"cbbd",
			"bb",
		},
		{
			"a",
			"a",
		},
		{
			"aa",
			"aa",
		},


	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, longestPalindrome(tc.input), "输入:%v", tc)
	}
}