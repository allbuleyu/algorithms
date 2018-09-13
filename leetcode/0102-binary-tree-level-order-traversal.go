package leetcode

/**
 * https://leetcode.com/problems/binary-tree-level-order-traversal/description/
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	
	var preOrder func(node *TreeNode, height int)
	
	preOrder = func(root *TreeNode, height int) {
		if root == nil {
			return
		}

		if height >= len(res) {
			res = append(res, []int{})
		}
		res[height] = append(res[height], root.Val)

		preOrder(root.Left, height+1)
		preOrder(root.Right, height+1)
	}

	preOrder(root, 0)
	
	return res
}

//var rmap map[int][]int
//func levelOrder(root *TreeNode) [][]int {
//	preOrder(root, 0)
//	res := make([][]int, len(rmap)+1)
//
//
//
//	for i := 0; i < len(rmap); i++ {
//		res = append(res, rmap[i])
//	}
//
//	return res
//}
//
//func preOrder(root *TreeNode, height int) {
//	if root == nil {
//		return
//	}
//
//	_, ok := rmap[height]
//	if !ok {
//		rmap[height] = make([]int, 0)
//	}
//	rmap[height] = append(rmap[height], root.Val)
//
//	preOrder(root.Left, height+1)
//	preOrder(root.Right, height+1)
//}

