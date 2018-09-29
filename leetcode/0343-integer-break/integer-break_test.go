package prob0343

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_prob0053(t *testing.T) {
	ast := assert.New(t)

	tcs := []struct {
		input int
		ans   int
	}{
		{2, 1},
		{3, 2},
		{10, 36},
		{8, 18},
		{9, 27},
	}

	for _, tc := range tcs {

		isEqual := ast.Equal(tc.ans, integerBreak(tc.input), "input is %v", tc.input)
		if isEqual {
			fmt.Printf("test case passed: ~~~%v~~~ \n", tc)
			fmt.Printf("%*s input: %v and expected is: %v \n", 10, " ", tc.input, tc.ans)
		}
	}
}