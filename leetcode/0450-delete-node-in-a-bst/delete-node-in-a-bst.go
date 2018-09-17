package prob0450


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
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	x := bstSearch(root, key)

	if x == nil {
		return root
	}

	if x.Left == nil {
		x = x.Right
	}else if x.Right == nil {
		x = x.Left
	}else {
		y := minimum(root.Right)
		if y != x.Right {
			y.Right = x.Right
			y.Left = x.Left
		}

		x = y
	}

	// 如果删除的不是root直接关联的子孩子,那么root引用不会有变化
	if x != root.Left && x != root.Right {
		return root
	}

	
	return root
}

func  bstSearch(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	var search *TreeNode

	if root.Val < key {
		search = bstSearch(root.Right, key)
	}else if root.Val > key {
		search = bstSearch(root.Left, key)
	}else {
		search = root
	}

	return search
}

func bstSearch2(root *TreeNode, key int) *TreeNode {
	search := root

	for search != nil {
		if search.Val < key {
			search = search.Right
		}else if search.Val > key {
			search = search.Left
		}else {
			return search
		}
	}
	
	return nil
}

// 找到二叉搜索树 key的后继 
func bstSuccessor(root *TreeNode, key int) *TreeNode {
	return nil
}

func minimum(root *TreeNode) *TreeNode {
	x := root
	if root.Left != nil {
		x = minimum(root.Left)
	}

	return x
}

func minimum1(root *TreeNode) *TreeNode {
	var x *TreeNode
	for {
		if root.Left == nil {
			break
		}

		x = root.Left
	}

	return x
}