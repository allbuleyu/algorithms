package prob0094

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
func inorderTraversal(root *TreeNode) []int {
	return helpIteration(root)
}

func helpIteration(root *TreeNode) []int {
	ans := make([]int, 0)
	stack := make([]*TreeNode, 0)
	node := root

	for len(stack) > 0 || node != nil {
		if node != nil {
			stack = append(stack, node)
			node = node.Left
		}else {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans = append(ans, node.Val)		// 与前序遍历不同的地方
			node = node.Right
		}
	}


	return ans
}

func helpMorrisTraversal(root *TreeNode) []int {
	cur := root
	ans := make([]int, 0)
	for cur != nil {
		if cur.Left != nil {
			precursor := cur.Left
			for precursor.Right != nil && precursor.Right == cur {
				precursor = precursor.Right
			}

			if precursor.Right == cur {
				ans = append(ans, cur.Val)
				precursor.Right = nil
				cur = cur.Right
			}else {
				precursor.Right = cur
				cur = cur.Left
			}

		}else {
			ans = append(ans, cur.Val)
			cur = cur.Right
		}
	}

	return ans
}