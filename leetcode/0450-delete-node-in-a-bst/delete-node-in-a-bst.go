package prob0450

import "github.com/allbuleyu/algorithms/kit"

//https://leetcode.com/problems/delete-node-in-a-bst/description/
//
//Given a root node reference of a BST and a key, delete the node with the given key in the BST. Return the root node reference (possibly updated) of the BST.
//
//Basically, the deletion can be divided into two stages:
//
//Search for a node to remove.
//If the node is found, delete the node.
//Note: Time complexity should be O(height of tree).
//
//Example:
//
//root = [5,3,6,2,4,null,7]
//key = 3
//
//5
/// \
//3   6
/// \   \
//2   4   7
//
//Given key to delete is 3. So we find the node with value 3 and delete it.
//
//One valid answer is [5,4,6,2,null,null,7], shown in the following BST.
//
//5
/// \
//4   6
///     \
//2       7
//
//Another valid answer is [5,2,6,null,4,null,7].
//
//5
/// \
//2   6
//\   \
//4   7

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode = kit.TreeNode

func deleteNode(root *TreeNode, key int) *TreeNode {
	return helpFunc(root, key)
}

func helpFunc(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > key {
		root.Left = helpFunc(root.Left, key)
	} else if root.Val < key {
		root.Right = helpFunc(root.Right, key)
	}

	if root.Left == nil {
		return root.Right
	} else if root.Right == nil {
		return root.Left
	}

	x := min(root.Right)
	x.Right = deleteMin(root.Right)
	x.Left = root.Left

	return x
}

func deleteMin(root *TreeNode) *TreeNode {
	if root.Left == nil {
		// right := root.Right
		// root.Right = nil
		// return right
		return root.Right
	}

	root.Left = deleteMin(root.Left)
	return root
}

func min(root *TreeNode) *TreeNode {
	if root.Left == nil {
		return root
	}

	return min(root.Left)
}

func deleteMinIteration(root *TreeNode) *TreeNode {
	s := make([]*TreeNode, 1)
	s[0] = root
	for len(s) > 0 {
		x := s[len(s)-1]
		if x.Left == nil {
			s[len(s)-1] = x.Right
			break
		}
		s = append(s, x.Left)
	}

	for i := 1; i < len(s); i++ {
		s[i-1].Left = s[i]
	}

	return s[0]
}
