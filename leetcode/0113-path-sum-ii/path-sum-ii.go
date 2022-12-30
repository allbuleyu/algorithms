package prob0113

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
func pathSum(root *TreeNode, sum int) [][]int {
	res := make([][]int, 0)
	cur := make([]int, 0)
	recursion(&res, &cur, root, sum)

	return res
}

// backtracking
func recursion(res *[][]int, cur *[]int, root *TreeNode, sum int) {
	if root == nil {
		return
	}

	sum -= root.Val
	*cur = append(*cur, root.Val)

	if root.Left == nil && root.Right == nil && sum == 0 {
		c := make([]int, len(*cur))
		copy(c, *cur)
		*res = append(*res, c)

		*cur = (*cur)[:len(*cur)-1]
		return
	}

	recursion(res, cur, root.Left, sum)
	recursion(res, cur, root.Right, sum)
	*cur = (*cur)[:len(*cur)-1]
}