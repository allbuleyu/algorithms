package prob0845

import (
	"testing"

	"fmt"
	"github.com/stretchr/testify/assert"
)

func Test_longestMountain(t *testing.T)  {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input []int
		ans int
	}{
		{[]int{0,1,2,3,4,5,6,7,8,9}, 0},
		{[]int{2,1,4,7,3,2,5}, 5},
		{[]int{2,2,2}, 0},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, longestMountain(tc.input), "输入:%v", tc)
	}
}