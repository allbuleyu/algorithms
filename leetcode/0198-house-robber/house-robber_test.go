package prob0198

import (
	"testing"
	"github.com/stretchr/testify/assert"

	"fmt"
)

func Test_prob0198(t *testing.T) {
	ast := assert.New(t)

	tcs := []struct{
		input []int
		ans int
	}{
		{[]int{1, 2, 3, 1}, 4},
		{[]int{2, 7, 1, 2, 9, 3, 1}, 17},
		{[]int{2, 3, 2}, 4},
		{[]int{2, 1}, 2},		// 在忽略基本情况的条件下,这个测试过不了
		{[]int{2, 1, 1, 2}, 4},
		{[]int{1},1},

	}

	for _, tc := range tcs {


		isEqual := ast.Equal(tc.ans, dynamicProgramming(tc.input), "input is %v", tc.input)
		if isEqual {
			fmt.Printf("test case passed: ~~~%v~~~ \n", tc)
			fmt.Printf("%*s input: %v and expected is: %v \n", 10," ", tc.input, tc.ans)
		}
	}
}