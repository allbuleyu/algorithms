package prob0014

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"fmt"
)

func Test_prb0014(t *testing.T)  {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input []string
		ans string
	}{
		{[]string{"flower","flow","flight"}, "fl",},
		{[]string{"dog","racecar","car"}, "",},
		{[]string{"dog"}, "dog",},
		{[]string{"flower","flow"}, "flow",},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, longestCommonPrefix(tc.input), "输入:%v", tc)
	}
}