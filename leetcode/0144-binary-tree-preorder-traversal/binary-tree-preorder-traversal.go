package prob0144

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
func preorderTraversal(root *TreeNode) []int {
	return preOrderIterate(root)
}

// 前序遍历
func preOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	*res = append(*res, root.Val)
	if root.Left != nil {
		preOrder(root.Left, res)
	}

	if root.Right != nil {
		preOrder(root.Right, res)
	}

}

// 迭代法前序遍历
func preOrderIterate(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	node := root

	for node != nil || len(stack) != 0 {
		for node != nil {
			res = append(res, node.Val)		// Add before going to children
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node = node.Right
	}

	return res
}