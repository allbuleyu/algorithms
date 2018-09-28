package prob0070

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_prob0070(t *testing.T) {
	ast := assert.New(t)

	tcs := []struct {
		input int
		ans   int
	}{
		// {0, 1}, 正整数,没有零
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{10, 89},
		{44, 1134903170},
	}

	for _, tc := range tcs {

		isEqual := ast.Equal(tc.ans, climbStairs(tc.input), "input is %v", tc.input)
		if isEqual {
			fmt.Printf("test case passed: ~~~%v~~~ \n", tc)
			fmt.Printf("%*s input: %v and expected is: %v \n", 10, " ", tc.input, tc.ans)
		}
	}
}