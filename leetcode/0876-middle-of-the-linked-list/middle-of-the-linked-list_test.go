package prob0876

import (
	"github.com/allbuleyu/algorithms/kit"
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
		{[]int{1, 2, 3, 4, 5, 6}, []int{4, 5, 6}},
		{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5}},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		actual := kit.TransferNodes(middleNode(kit.CreateNodes(tc.input)))
		ast.Equal(tc.ans, actual, "输入:%v", tc)
	}
}