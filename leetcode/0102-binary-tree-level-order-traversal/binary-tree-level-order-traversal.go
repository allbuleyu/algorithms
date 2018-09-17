package prob0102


/*
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

func PreOrder(root *TreeNode) [][]int {
	return levelOrder(root)
}

//var rmap map[int][]int
//func levelOrder(root *TreeNode) [][]int {
//	rmap = make(map[int][]int)
//	preOrder(root, 0)
//	res := make([][]int, 0)
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
//		rmap[height] = []int{}
//	}
//
//	rmap[height] = append(rmap[height], root.Val)
//
//	preOrder(root.Left, height+1)
//	preOrder(root.Right, height+1)
//}

