package kit

import "sync"

type TreeInts []int

const Null int = -1 << 63

type orderLevel int
const (
	_  orderLevel = iota
	LevelPreOrder
	LevelInOrder
	LevelPostOrder
)

type TreeInterface interface {
	PreOrder() []int
	InOrder() []int
	PostOrder() []int

	String() string
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type Trees struct {
	Head *TreeNode
	depth int
	sync.RWMutex
}

func NewTrees(ints []int) Trees {
	n := len(ints)
	if n == 0 {
		panic("tree arr can not be zero")
	}

	root := &TreeNode{
		Val: ints[0],
	}


	
	return Trees{
		Head:root,
		depth:0,
		RWMutex:sync.RWMutex{},
	}
}

func (t *Trees) PreOrder() []int {
	res := make([]int, 0)
	toIntsPreOrder(t.Head, &res)
	return res
}

func (t *Trees) InOrder() []int {
	res := make([]int, 0)
	toIntsInOrder(t.Head, &res)
	return res
}

func (t *Trees) PostOrder() []int {
	res := make([]int, 0)
	toIntsPostOrder(t.Head, &res)
	return res
}

func (t *Trees) String() string {

	return ""
}

func toIntsPreOrder(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}

	*res = append(*res, node.Val)
	if node.Left != nil {
		toIntsPreOrder(node.Left, res)
	}

	if node.Right != nil {
		toIntsPreOrder(node.Right, res)
	}

}

func toIntsInOrder(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}

	if node.Left != nil {
		toIntsInOrder(node.Left, res)
	}
	*res = append(*res, node.Val)

	if node.Right != nil {
		toIntsInOrder(node.Right, res)
	}

}

func toIntsPostOrder(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}

	if node.Left != nil {
		toIntsPostOrder(node.Left, res)
	}

	if node.Right != nil {
		toIntsPostOrder(node.Right, res)
	}
	*res = append(*res, node.Val)
}

