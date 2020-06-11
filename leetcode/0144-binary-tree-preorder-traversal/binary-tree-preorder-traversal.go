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
	return helpIterate(root)
}

func helpIterate(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	node := root

	for len(stack) > 0 || node != nil {
		if node != nil {
			stack = append(stack, node)
			res = append(res, node.Val)
			node = node.Left
		}else {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			node = node.Right
		}
	}

	return res
}

func helpMorrisTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	cur := root
	for cur != nil {

		if cur.Left != nil {
			precursor := cur.Left
			for precursor.Right != nil && precursor.Right == cur {
				precursor = precursor.Right
			}

			if precursor.Right == cur {

				precursor.Right = nil
				cur = cur.Right
			}else {
				ans = append(ans, cur.Val)
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