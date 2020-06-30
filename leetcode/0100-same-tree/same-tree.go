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
	return helpIteration(p,q)
}

func helpRecursion(p,q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return helpRecursion(p.Left, q.Left) && helpRecursion(p.Right, q.Right)
}

func helpIteration(p,q *TreeNode) bool {
	sp := make([]*TreeNode, 0)
	sq := make([]*TreeNode, 0)

	for p != nil || q != nil {
		if p == nil || q == nil {
			return false
		}

		if p.Val != q.Val {
			return false
		}

		if p.Left != nil || q.Left != nil {
			sp = append(sp, p.Left)
			sq = append(sq, q.Left)
		}

		if p.Right != nil || q.Right != nil {
			sp = append(sp, p.Right)
			sq = append(sq, q.Right)
		}

		if len(sp) == 0 && 0 == len(sq) {
			return true
		}

		p = sp[len(sp)-1]
		sp = sp[:len(sp)-1]
		q = sq[len(sq)-1]
		sq = sq[:len(sq)-1]
	}

	return true
}