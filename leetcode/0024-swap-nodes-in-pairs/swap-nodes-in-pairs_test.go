package prob0024

import (
	"algorithms/kit"
	"testing"

	"github.com/stretchr/testify/assert"
	"fmt"
)

func Test_swapPairs(t *testing.T)  {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input []int
		ans []int
	}{
		{[]int{1,2,3,4,5}, []int{2,1,4,3,5}},
		{[]int{1,2,3,4}, []int{2,1,4,3}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},

	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		actual := kit.TransferNodes(swapPairs(kit.CreateNodes(tc.input)))
		ast.Equal(tc.ans, actual, "输入:%v", tc)
	}
}