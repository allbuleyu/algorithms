package prob0674

import (
	"testing"

	"fmt"
	"github.com/stretchr/testify/assert"
)

func Test_findLengthOfLCIS(t *testing.T)  {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input []int

		ans int
	}{
		{[]int{1,3,5,4,7}, 3},
		{[]int{1,3,5,6,7}, 5},
		{[]int{2,2,2,2}, 1},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, findLengthOfLCIS(tc.input), "输入:%v", tc)
	}
}