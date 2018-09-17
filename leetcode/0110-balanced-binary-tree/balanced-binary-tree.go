package prob0110


/**
 * https://leetcode.com/problems/balanced-binary-tree/description/
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	_, isBalanced := findDepth(root)

	return isBalanced
}

func findDepth(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	ldepth, lbool := findDepth(root.Left)
	rdepth, rbool := findDepth(root.Right)

	if !lbool || !rbool || sbuAbs(ldepth, rdepth) > 1 {
		return 0, false
	}

	return max(ldepth, rdepth) + 1, true
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func sbuAbs(a, b int) int {
	if a > b {
		return a-b
	}

	return b-a
}