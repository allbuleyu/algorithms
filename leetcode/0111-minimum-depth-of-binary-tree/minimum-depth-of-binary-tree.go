package prob0111



/**
 * https://leetcode.com/problems/minimum-depth-of-binary-tree/description/
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Right != nil && root.Left == nil {
		return minDepth(root.Right) + 1
	}

	if root.Left != nil && root.Right == nil {
		return minDepth(root.Left) + 1
	}

	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}