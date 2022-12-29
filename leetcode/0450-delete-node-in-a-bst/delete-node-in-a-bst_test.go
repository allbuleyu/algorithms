package prob0450

import (
	"fmt"
	"testing"

	"algorithms/kit"
	"github.com/stretchr/testify/assert"
)

func Test_deleteNode(t *testing.T) {
	type args struct {
		root *TreeNode
		key  int
	}

	ast := assert.New(t)

	// test case
	tcs := []struct {
		input  []int
		input2 int
		ans    []int
	}{
		{[]int{5, 3, 6, 2, 4, kit.Null, 7}, 3, []int{5, 4, 6, 2, kit.Null, kit.Null, 7}},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		tree := kit.NewTrees(tc.input)
		actual := kit.InOrder(deleteNode(tree.Root, tc.input2))
		ast.Equal(tc.ans, actual, "输入:%v", tc)
	}
}
