package prob0144

//https://leetcode.com/problems/binary-tree-preorder-traversal/description/
//
//Given a binary tree, return the preorder traversal of its nodes' values.
//
//Example:
//
//Input: [1,null,2,3]
//1
//\
//2
///
//3
//
//Output: [1,2,3]
//Follow up: Recursive solution is trivial, could you do it iteratively?

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	rightStack := make([]*TreeNode, 0)
	res := make([]int, 0)

	for cur := root; cur != nil; {
		res = append(res, cur.Val)

		if cur.Left != nil {
			if cur.Right != nil {
				rightStack = append(rightStack, cur.Right)
			}

			cur = cur.Left
		}else if cur.Right != nil {
			cur = cur.Right
		}else {

			if len(rightStack) == 0 {
				break
			}

			cur = rightStack[len(rightStack)-1]
			rightStack = rightStack[:len(rightStack)-1]
		}

	}

	return res
}

//sample 0 ms submission
//func preorderTraversal(root *TreeNode) []int {
//	if root == nil {
//		return []int{}
//	}
//
//	nodes, stack := []int{}, []*TreeNode{root}
//
//	for len(stack) > 0 {
//
//		if root != nil {
//			nodes = append(nodes, root.Val)
//
//			if root.Right != nil {
//				stack = append(stack, root.Right)
//			}
//
//			root = root.Left
//		} else {
//			root = stack[len(stack)-1]
//			if len(stack) > 0 {
//				stack = stack[:len(stack)-1]
//			}
//		}
//	}
//	return nodes
//
//}