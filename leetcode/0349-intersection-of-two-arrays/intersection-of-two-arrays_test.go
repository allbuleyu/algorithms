package prob0349

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_prob0349(t *testing.T) {
	ast := assert.New(t)

	tcs := []struct {
		input, input1 []int
		ans   []int
	}{
		{[]int{1,2,2,1}, []int{2,2}, []int{2}},
		{[]int{4,9,5}, []int{9,4,9,8,4}, []int{4,9}},
	}

	for _, tc := range tcs {

		isEqual := ast.Equal(tc.ans, intersection(tc.input, tc.input1), "input is %v", tc.input)
		if isEqual {
			fmt.Printf("test case passed: ~~~%v~~~ \n", tc)
			fmt.Printf("%*s input: %v and expected is: %v \n", 10, " ", tc.input, tc.ans)
		}
	}
}