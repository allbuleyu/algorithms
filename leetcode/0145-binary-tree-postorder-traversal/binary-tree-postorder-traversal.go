package prob0145

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
func postorderTraversal(root *TreeNode) []int {
	return helpIteration1(root)
}

func helpMirrorTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	node := root
	for node != nil {
		if node.Right != nil {
			precursor := node.Right
			for precursor.Left != nil && node != precursor.Left {
				precursor = precursor.Left
			}

			if node == precursor.Left {
				precursor.Left = nil
				node = node.Left
			}else {
				ans = append(ans, node.Val)
				precursor.Left = node
				node = node.Right
			}
		}else {
			ans = append(ans, node.Val)
			node = node.Left
		}
	}

	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}

	return ans
}


func helpIteration2(root *TreeNode) []int {
	ans := make([]int, 0)
	stack := make([]*TreeNode, 0)
	node := root
	for len(stack) > 0 || node != nil {
		if node.Right != nil {
			ans = append(ans, node.Val)

			stack = append(stack, node)
			node = node.Right
		}else {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			node = node.Left
		}
	}

	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}

	return ans
}

func helpIteration1(root *TreeNode) []int {
	ans := make([]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	head := root
	for len(stack) > 0 {
		t := stack[len(stack)-1]
		if (t.Left != nil && t.Right != nil) || t.Left == head || t.Right == head {
			ans = append(ans, t.Val)
			stack = stack[:len(stack)-1]
			head = t
		}else {
			if t.Right != nil {
				stack = append(stack, t.Right)
			}

			if t.Left != nil {
				stack = append(stack, t.Left)
			}
		}
	}

	return ans
}


func helpRecursionTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	ans := make([]int, 0)

	ans = append(ans, helpRecursionTraversal(root.Left)...)
	ans = append(ans, helpRecursionTraversal(root.Left)...)
	ans = append(ans, root.Val)

	return ans
}