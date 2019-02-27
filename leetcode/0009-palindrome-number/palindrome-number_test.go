package prob0009

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"fmt"
)

func Test_prb0001(t *testing.T)  {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input int
		ans bool
	}{
		{121, true,},
		{1, true,},
		{-1, false,},
		{10, false,},

	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, isPalindrome(tc.input), "输入:%v", tc)
	}
}