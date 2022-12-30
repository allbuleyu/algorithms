package prob0235

import (
	"fmt"
	"algorithms/kit"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input []int
		target int
		ans []int
	}{
		{[]int{5,3,6,2,4,kit.Null,7}, 3, []int{2, 4, 5, 6, 7}},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		root := kit.NewTrees(tc.input)
		ans := deleteNode(root.Root, tc.target)
		ast.Equal(tc.ans, kit.InOrder(ans), "输入:%v", tc)
	}
}