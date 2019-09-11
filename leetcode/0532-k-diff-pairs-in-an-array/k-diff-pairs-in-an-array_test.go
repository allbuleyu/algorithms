package prob0532

import (
	"testing"

	"fmt"
	"github.com/stretchr/testify/assert"
)

func Test_findPairs(t *testing.T)  {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input []int
		k int
		ans int
	}{
		{[]int{3, 1, 4, 1, 5},2,2},
		{[]int{1, 2, 3, 4, 5},1,4},
		{[]int{1, 3, 1, 5, 4},0,1},


	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, findPairs(tc.input, tc.k), "输入:%v", tc)
	}
}