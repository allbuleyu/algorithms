package prob0236

import "algorithms/kit"

type TreeNode = kit.TreeNode
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	ans := &TreeNode{}
	helpRecursion(root,p,q,ans)
	return ans.Left
}

func helpRecursion(root,p,q, ans *TreeNode) bool {
	if root == nil {
		return false
	}

	if ans.Left != nil {
		return false
	}

	var self bool
	if root.Val == p.Val || root.Val == q.Val {
		self = true
	}

	cl := helpRecursion(root.Left, p, q, ans)
	cr := helpRecursion(root.Right, p, q, ans)

	if (self && cl) || (self && cr) || (cl && cr) {
		ans.Left = root
	}

	return self || cl || cr
}