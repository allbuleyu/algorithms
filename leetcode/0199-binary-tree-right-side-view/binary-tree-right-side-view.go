package prob0199

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
func rightSideView(root *TreeNode) []int {
	return dfsIterate(root)
}

func dfsIterate(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	m := make(map[int]int)
	nodeStack := make([]*TreeNode,0)
	depthStack := make([]int, 0)

	nodeStack = append(nodeStack, root)
	depthStack = append(depthStack, 0)
	for len(nodeStack) != 0 {
		node := nodeStack[len(nodeStack)-1]
		depth := depthStack[len(depthStack)-1]

		nodeStack = nodeStack[:len(nodeStack)-1]
		depthStack = depthStack[:len(depthStack)-1]

		if node != nil {
			_, ok := m[depth]
			if !ok {
				m[depth] = node.Val
			}


			nodeStack = append(nodeStack, node.Left, node.Right)
			depthStack = append(depthStack, depth+1, depth+1)
		}

	}

	res := make([]int, len(m))
	for i := 0; i < len(m); i++ {
		res[i] = m[i]
	}

	return res
}

func bfsIterate(root *TreeNode) []int {
	return nil
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}