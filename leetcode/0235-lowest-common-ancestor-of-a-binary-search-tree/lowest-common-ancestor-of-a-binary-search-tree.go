package prob0235

import "github.com/allbuleyu/algorithms/kit"

type TreeNode = kit.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	ans := &TreeNode{}
	helper(root, p, q, ans)
	return ans.Right
}

func helper(root, p, q, ans *TreeNode) bool {
	if root == nil {
		return false
	}

	cur := false
	if root == p || root == q {
		cur = true
	}

	left := helper(root.Left, p, q, ans)
	right := helper(root.Right, p, q, ans)

	if (cur && left) || (cur && right) {
		ans.Right = root
		return false
	}

	return left || right || cur
}