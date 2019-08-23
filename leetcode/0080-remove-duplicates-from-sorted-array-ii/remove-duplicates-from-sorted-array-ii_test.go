package prob0080

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"fmt"
)

func Test_removeDuplicates(t *testing.T)  {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input []int
		ans []int
	}{
		{[]int{1,1,1,2,2,3}, []int{1,1,2,2,3},},
		{[]int{0,0,1,1,1,1,2,3,3}, []int{0,0,1,1,2,3,3}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, tc.input[:removeDuplicates(tc.input)], "输入:%v", tc)
	}
}