package kit

import (
	"fmt"
	"sync"
)

type BinarySearchTree struct {
	root *TreeNode
	depth int
	lock sync.RWMutex
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
	if node.Val < root.Val {
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

func (bst *BinarySearchTree) InOrder() []int {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	return inOrderBST(bst.root)
}

func inOrderBST(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)
	res = append(res, inOrderBST(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inOrderBST(root.Right)...)

	return res
}

func (bst *BinarySearchTree) PreOrder() []int {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	return preOrderBST(bst.root)
}

func PreOrderBST(root *TreeNode) []int {
	return preOrderBST(root)
}

func preOrderBST(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)
	res = append(res, root.Val)
	res = append(res, preOrderBST(root.Left)...)
	res = append(res, preOrderBST(root.Right)...)

	return res
}

func (bst *BinarySearchTree) PostOrder() []int {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	return postOrderBST(bst.root)
}

func postOrderBST(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)
	res = append(res, postOrderBST(root.Left)...)
	res = append(res, postOrderBST(root.Right)...)
	res = append(res, root.Val)

	return res
}

func (bst *BinarySearchTree) Max() int {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	return bstMax(bst.root)
}

func bstMax(root *TreeNode) int {
	if root == nil {
		return Null
	}

	if root.Right == nil {
		return root.Val
	}

	return bstMax(root.Right)
}

func (bst *BinarySearchTree) Min() int {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	return bstMin(bst.root)
}

func bstMin(root *TreeNode) int {
	if root == nil {
		return Null
	}

	if root.Left == nil {
		return root.Val
	}

	return bstMin(root.Left)
}

func (bst *BinarySearchTree) String()  {
	fmt.Println(" ---------------------------------------------------- ")
	stringify(bst.root, 0)
	fmt.Println(" ---------------------------------------------------- ")
}

func stringify(n *TreeNode, level int)  {
	if n != nil{
		format := ""
		for i:=0; i<level; i++{
			format += "              "
		}

		format += "-----["
		level ++
		stringify(n.Right, level)
		fmt.Printf(format+"%d\n", n.Val)
		stringify(n.Left, level)
	}
}
