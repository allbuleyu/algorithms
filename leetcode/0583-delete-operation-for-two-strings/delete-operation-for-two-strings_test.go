package prob0583

import (
	"github.com/stretchr/testify/assert"
	"fmt"

	"testing"
)

func Test_prob0583(t *testing.T)  {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		s1, s2 string
		ans int
	}{
		{
			"eat", "ate",
			2,
		},
		{
			"leet", "delete",
			4,
		},{
			"leetcode", "delete",
			6,
		},{
			"leetcode", "",
			8,
		},{
			"", "leetcode",
			8,
		},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, minDistance(tc.s1, tc.s2), "输入:%v", tc)
	}
}
