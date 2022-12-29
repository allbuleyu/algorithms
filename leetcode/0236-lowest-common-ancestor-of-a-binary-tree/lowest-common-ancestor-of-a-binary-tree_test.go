package prob0236

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
		p,q *TreeNode
		ans int
	}{
		{[]int{3,5,1,6,2,0,8,kit.Null,kit.Null,7,4}, &TreeNode{Val: 5}, &TreeNode{Val: 1}, 3},
		{[]int{3,5,1,6,2,0,8,kit.Null,kit.Null,7,4}, &TreeNode{Val: 5}, &TreeNode{Val: 4}, 5},

	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ans := lowestCommonAncestor(kit.NewTrees(tc.input).Root, tc.p, tc.q).Val
		ast.Equal(ans, tc.ans, "输入:%v", tc)
	}
}