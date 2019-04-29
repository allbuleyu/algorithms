package prob0094

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
func inorderTraversal(root *TreeNode) []int {
	return helpIterate(root)
}

func helpRecursion(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)

	res = append(res, inorderTraversal(root.Left)...)

	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)

	return res
}

// 迭代法中序遍历
func helpIterate(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	node := root

	for len(stack) > 0 || node != nil {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}


		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		node = node.Right
	}

	return res
}