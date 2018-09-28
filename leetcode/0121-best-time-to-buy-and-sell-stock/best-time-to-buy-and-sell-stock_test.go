package prob0121

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_prob0121(t *testing.T) {
	ast := assert.New(t)

	tcs := []struct {
		input []int
		ans   int
	}{
		{[]int{7,1,5,3,6,4}, 5},
		{[]int{7,6,4,3,1}, 0},
		{[]int{1, 2, 4}, 3},
		{[]int{1, 2}, 1},
		{[]int{2, 1}, 0},
		{[]int{1}, 0},
	}

	for _, tc := range tcs {

		isEqual := ast.Equal(tc.ans, maxProfit(tc.input), "input is %v", tc.input)
		if isEqual {
			fmt.Printf("test case passed: ~~~%v~~~ \n", tc)
			fmt.Printf("%*s input: %v and expected is: %v \n", 10, " ", tc.input, tc.ans)
		}
	}
}