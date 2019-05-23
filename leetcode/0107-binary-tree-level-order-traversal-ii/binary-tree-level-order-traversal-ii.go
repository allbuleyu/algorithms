package prob0107

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
func levelOrderBottom(root *TreeNode) [][]int {
	return iterate(root)
}

func recursion(root *TreeNode) [][]int {
	return nil
}

// bfs iterate solution
func iterate(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	stackDepth := make([]int, 0)
	stackDepth = append(stackDepth, 0)

	for len(stack) != 0 {
		node := stack[0]
		stack = stack[1:]
		depth := stackDepth[0]
		stackDepth = stackDepth[1:]

		if node == nil {
			continue
		}

		if len(res) <= depth {
			res = append(res, make([]int, 0))
		}

		res[depth] = append(res[depth], node.Val)

		stack = append(stack, node.Left, node.Right)
		stackDepth = append(stackDepth, depth+1, depth+1)
	}

	for i,j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}