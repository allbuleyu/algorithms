package prob0226

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
func invertTree(root *TreeNode) *TreeNode {
	return iterate(root)
}

func iterate(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			continue
		}

		node.Left, node.Right = node.Right, node.Left

		queue = append(queue, node.Left, node.Right)
	}

	return root
}