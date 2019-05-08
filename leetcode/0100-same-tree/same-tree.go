package prob0100

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
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	return isEqual(p, q) && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func isEqual(p, q *TreeNode) bool {
	if p.Val != q.Val {
		return false
	}
	
	return true
}