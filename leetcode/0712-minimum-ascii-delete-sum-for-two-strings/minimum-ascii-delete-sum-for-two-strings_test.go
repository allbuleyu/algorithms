package prob0712

import (
	"github.com/stretchr/testify/assert"
	"fmt"

	"testing"
)

func Test_prob0712(t *testing.T)  {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		s1, s2 string
		ans int
	}{
		{
			"eat", "ate",
			202,
		},
		{
			"sea", "eat",
			231,
		},
		{
			"delete", "leet",
			403,
		},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, minimumDeleteSum(tc.s1, tc.s2), "输入:%v", tc)
	}
}
