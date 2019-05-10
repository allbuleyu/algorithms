package prob0129

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
func sumNumbers(root *TreeNode) int {
	return dfsIterate(root)
}

func dfsRecursion(root *TreeNode) int {
	curSUm := 0
	return dfsRecursionHelp(root, curSUm)
}

func dfsRecursionHelp(root *TreeNode, curSum int) int {
	if root == nil {
		return 0
	}

	var res int

	curSum = curSum * 10 + root.Val
	if root.Left == nil && root.Right == nil {
		res = curSum
	}

	res += dfsRecursionHelp(root.Left, curSum)
	res += dfsRecursionHelp(root.Right, curSum)

	return res
}

func dfsIterate(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var res int
	stack := make([]*TreeNode, 0)
	sumStack := make([]int, 0)
	stack = append(stack, root)
	sumStack = append(sumStack, 0)

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		curSum := sumStack[len(sumStack)-1]

		stack = stack[:len(stack)-1]
		sumStack = sumStack[:len(sumStack)-1]

		if node == nil {
			continue
		}

		if node.Left == nil && node.Right == nil {
			res += curSum * 10 + node.Val
		}


		stack = append(stack, node.Left, node.Right)
		sumStack = append(sumStack, curSum * 10 + node.Val, curSum * 10 + node.Val)
	}

	return res
}