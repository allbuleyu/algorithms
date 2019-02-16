package prob0264

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
		ans int
	}{
		{
			10,
			12,
		},
		{
			1, 1,
		},
		{
			2, 2,
		},
		{
			7, 8,
		},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, nthUglyNumber(tc.input), "输入:%v", tc)
	}
}