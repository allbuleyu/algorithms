package prob0437

import "algorithms/kit"

type TreeNode = kit.TreeNode
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	return recursion(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)	// 这一步,很关键!
}

func recursion(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	res := 0
	sum -= root.Val
	if sum == 0 {
		res++
	}

	res += recursion(root.Left, sum)
	res += recursion(root.Right, sum)

	return res
}