package prob0083

import (
	"github.com/allbuleyu/algorithms/kit"
	"testing"

	"github.com/stretchr/testify/assert"
	"fmt"
)

func Test_prb0083(t *testing.T)  {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input []int
		ans []int
	}{
		{[]int{1, 2, 2}, []int{1, 2}},
		{[]int{1,1, 2, 6}, []int{1, 2, 6}},
		{[]int{1}, []int{1}},
		{[]int{1,1,1,1}, []int{1}},

	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		actual := kit.TransferNodes(deleteDuplicates(kit.CreateNodes(tc.input)))
		ast.Equal(tc.ans, actual, "输入:%v", tc)
	}
}