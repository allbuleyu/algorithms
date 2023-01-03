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
	return recursion(1, n)
}

func recursion(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	nodes := make([]*TreeNode, 0)
	for i := start; i < end; i++ {
		left, right := recursion(start, i-1), recursion(i+1, end)
		for _, l := range left {
			for _, r := range right {
				node := &TreeNode{Val: i, Left: l, Right: r}
				nodes = append(nodes, node)

			}
		}

	}

	return nodes
}
