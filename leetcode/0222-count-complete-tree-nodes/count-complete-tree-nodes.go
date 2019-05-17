package prob0222

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
func countNodes(root *TreeNode) int {
	d := findDepth(root)

	if d < 0 {
		return 0
	}

	if findDepth(root.Right) == d-1 {
		d = 1<<uint(d)
	}else {
		d = 1<<uint(d-1)
	}

	return d
}

func findDepth(root *TreeNode) int {
	if root == nil {
		return -1
	}

	depth := 0
	for root != nil {
		depth++
		root=root.Left
	}

	return depth
}