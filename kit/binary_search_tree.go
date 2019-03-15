package kit

import "sync"

type BinarySearchTree struct {
	root *TreeNode
	depth int
	lock sync.RWMutex
}

func (bst *BinarySearchTree) InOrder(nums []int) []int {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	return nums
}

func inOrderBST(root *TreeNode, nums []int) {

}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}

func (bst *BinarySearchTree) Insert(val int) {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	node := &TreeNode{Val:val}
	if bst.root == nil {
		bst.root = node
	}else {
		insertBSTNode(bst.root, node)
	}
}

func insertBSTNode(root, node *TreeNode) {
	if root.Val < node.Val {
		if root.Left == nil {
			root.Left = node
		}else {
			insertBSTNode(root.Left, node)
		}
	}else {
		if root.Right == nil {
			root.Right = node
		}else {
			insertBSTNode(root.Right, node)
		}
	}
}






