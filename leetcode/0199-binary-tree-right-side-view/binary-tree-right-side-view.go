package prob0199

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
func rightSideView(root *TreeNode) []int {
	return bfsIterate(root)
}

func dfsIterate(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	m := make(map[int]int)			// 用来判断和保存每一层最右边的节点
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


// 广度优先使用一个队列来当辅助
func bfsIterate(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	nodeQueue := make([]*TreeNode, 0)
	depthQueue := make([]int,0)
	nodeQueue = append(nodeQueue,root)
	depthQueue = append(depthQueue,0)

	m := make(map[int]int)
	max_depth := 0

	for len(nodeQueue) != 0 {
		node := nodeQueue[0]
		nodeQueue = nodeQueue[1:]

		depth := depthQueue[0]
		max_depth = max(depth, max_depth)
		depthQueue = depthQueue[1:]

		if node == nil {
			continue
		}

		if _, ok := m[max_depth]; !ok {
			m[max_depth] = node.Val
		}

		nodeQueue = append(nodeQueue, node.Right, node.Left)
		depthQueue = append(depthQueue, depth+1, depth+1)
	}

	res := make([]int, len(m))
	for i := 0; i < len(m); i++ {
		res[i] = m[i]
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}