package prob0035

import (
	"testing"

	"fmt"
	"github.com/stretchr/testify/assert"
)

func Test_prb0014(t *testing.T)  {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input []int
		target int
		ans int
	}{
		{[]int{1,3,5,6}, 5, 2},
		{[]int{1,3,5,6}, 2, 1},
		{[]int{1,3,5,6}, 7, 4},
		{[]int{1,3,5,6}, 0, 0},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, searchInsert(tc.input, tc.target), "输入:%v", tc)
	}
}