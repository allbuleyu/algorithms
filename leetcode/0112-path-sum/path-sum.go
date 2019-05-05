package prob0112

import "github.com/allbuleyu/algorithms/kit"

type TreeNode = kit.TreeNode
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, sum int) bool {
	return recursion(root,sum)
}

func recursion(root *TreeNode, sum int) bool {
	res := false
	recursionHelp(root, &sum, &res)

	return res
}

func recursionHelp(root *TreeNode, sum *int, res *bool) {
	if root.Left == nil && root.Right == nil {
		if *sum == 0 {
			*res = true
		}
	}

	if root.Left != nil {
		*sum -= root.Val
		recursionHelp(root.Left, sum, res)
		*sum += root.Val
	}

	if root.Right != nil {
		*sum -= root.Val
		recursionHelp(root.Right, sum, res)
		*sum += root.Val
	}
}