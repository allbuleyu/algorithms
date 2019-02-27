package prob0026

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"fmt"
)

func Test_prb0014(t *testing.T)  {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input []int
		ans []int
	}{
		{[]int{1,1,2}, []int{1,2},},
		{[]int{0,1,1,1,2,2,3,3,4}, []int{0,1,2,3,4}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, tc.input[:removeDuplicates(tc.input)], "输入:%v", tc)
	}
}