package prob0881

import (
	"testing"

	"fmt"
	"github.com/stretchr/testify/assert"
)

func Test_numRescueBoats(t *testing.T)  {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input []int
		k int
		ans int
	}{
		//{[]int{10, 5, 2, 6},100,8},
		{[]int{1,2,3}, 3, 2},
		{[]int{3,2,2,1}, 3, 3},
		{[]int{3,2,2,1}, 5, 2},
		{[]int{3,5,3,4}, 5, 4},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, numRescueBoats(tc.input, tc.k), "输入:%v", tc)
	}
}