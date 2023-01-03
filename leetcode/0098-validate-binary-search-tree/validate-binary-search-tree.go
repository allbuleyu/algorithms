package prob0098

import (
	"algorithms/kit"
	"math"
)

type TreeNode = kit.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	return morris(root)
}

// [5,1,4,null,null,3,6]
// [5,1,6,null,null,3]
func help(root, low, high *TreeNode) bool {
	if root == nil {
		return true
	}

	if (low != nil && root.Val <= low.Val) || (high != nil && root.Val >= high.Val) {
		return false
	}

	return help(root.Left, low, root) && help(root.Right, root, high)
}

func morris(root *TreeNode) bool {
	pv := math.MinInt64
	for root != nil {
		if root.Left != nil {
			prev := root.Left
			for prev.Right != nil && prev.Right != root {
				prev = prev.Right
			}

			if prev.Right == nil {
				prev.Right = root

				root = root.Left
			} else {
				prev.Right = nil

				if root.Val <= pv {
					return false
				}

				pv = root.Val
				root = root.Right
			}
		} else {
			if root.Val <= pv {
				return false
			}
			pv = root.Val
			root = root.Right
		}
	}

	return true
}
