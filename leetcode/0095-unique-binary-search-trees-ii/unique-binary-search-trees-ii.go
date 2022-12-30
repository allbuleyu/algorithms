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
	return recursion(0, n)
}

func recursion(start, end int) []*TreeNode {
	nodes := make([]*TreeNode, 0)
	for i := start; i <= end; i++ {
		left := recursion(start, i-1)
		right := recursion(i+1, end)

		node := &TreeNode{Val: i}

		for _, l := range left {
			for _, r := range right {
				node.Left = l
				node.Right = r

				nodes = append(nodes, node)
			}
		}
	}

	return nodes
}
