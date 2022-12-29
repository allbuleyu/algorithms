package prob0095

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
func generateTrees(n int) []*TreeNode {
	return help(0, n)
}

func help(start, end int) []*TreeNode {
	if start > end {
		return nil
	}

	nodes := make([]*TreeNode, 0)
	for i := start; i <= end; i++ {
		left := help(start, i-1)
		right := help(i+1, end)

		node := &TreeNode{Val: i}

		if len(left) == 0 && len(right) == 0 {
			nodes = append(nodes, node)
			return nodes
		}

	}

	return nodes
}
