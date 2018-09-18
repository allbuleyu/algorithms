package prob0007


import (
	"testing"

	"github.com/stretchr/testify/assert"
	"fmt"
)

func Test_prb0007(t *testing.T)  {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input, ans int
	}{
		{
			12345,
			54321,
		},
		{
			123456,
			654321,
		},
		{
			123457,
			754321,
		},
		{
			2147483647,
			0,
		},
		{
			1534236469,
			0,
		},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, reverse(tc.input), "输入:%v", tc)
	}
}