package prob0342

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)
var tcs = []struct{
	input int
	ans bool
}{
	{8, false},
	{4, true},
	{2, false},
	{1, true},
	{0, false},
	{5, false},

}

func Test_isPowerOfFour(t *testing.T) {
	ast := assert.New(t)

	// test case
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, isPowerOfFour(tc.input), "输入:%v", tc)
	}
}