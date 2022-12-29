package prob0106

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

func buildTree(inorder []int, postorder []int) *TreeNode {
	return helpRecursion(inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1)
}

// 用hash map 速度会快很多
func helpRecursion(inorder []int, iLeft, iRight int, postorder []int, pLeft, pRight int) *TreeNode {
	if iLeft > iRight || pLeft > pRight {
		return nil
	}

	root := &TreeNode{Val:postorder[pRight]}
	i := 0
	for i = iLeft; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			break
		}
	}

	root.Left = helpRecursion(inorder, iLeft, i-1, postorder, pLeft, pLeft+(i-iLeft)-1)
	root.Right = helpRecursion(inorder, i+1, iRight, postorder, pLeft+(i-iLeft), pRight-1)

	return root
}