package prob0101

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
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isMirrorIterate(root)
}

func isMirror(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && isMirror(p.Left, q.Right) && isMirror(p.Right, q.Left)
}

func isMirrorIterate(root *TreeNode) bool {
	stack := make([]*TreeNode, 0)
	stack = append(stack, root, root)
	for len(stack) != 0 {
		l, r := stack[len(stack)-2], stack[len(stack)-1]
		stack = stack[:len(stack)-2]
		if l == nil && r == nil {
			continue
		}

		if l == nil || r == nil {
			return false
		}

		if l.Val != r.Val {
			return false
		}

		stack = append(stack,l.Left, r.Right, l.Right, r.Left)
	}

	return true
}