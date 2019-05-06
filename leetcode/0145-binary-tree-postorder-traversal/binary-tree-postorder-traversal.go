package prob0145

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
func postorderTraversal(root *TreeNode) []int {
	return postOrderIterate(root)
}

func postOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	if root.Left != nil {
		postOrder(root.Left, res)
	}

	if root.Right != nil {
		postOrder(root.Right, res)
	}
	*res = append(*res, root.Val)
}

// 迭代法,后续遍历
func postOrderIterate(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)

	stack := make([]*TreeNode, 0)
	node := root
	for node != nil || len(stack) != 0 {
		for node != nil {
			res = append(res, node.Val)		// Add after all left children
			stack = append(stack, node)
			node = node.Right
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node = node.Left
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1,j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}