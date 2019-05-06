package kit

import (
	"bytes"
	"fmt"
	"sync"
)

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
	Root *TreeNode
	depth int
	sync.RWMutex
}

func NewTrees(ints []int) Trees {
	n := len(ints)
	if n == 0 {
		panic("tree array can not be zero")
	}

	root := &TreeNode{
		Val: ints[0],
	}

	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root

	isDepth := false
	depth := 0

	for i := 1; i < n; {
		root := queue[0]
		queue = queue[1:]

		if ints[i] != Null {
			root.Left = &TreeNode{Val:ints[i]}
			queue = append(queue, root.Left)

			if !isDepth {
				depth++
				isDepth = true
			}
		}
		i++

		if i < n && ints[i] != Null {
			root.Right = &TreeNode{Val:ints[i]}
			queue = append(queue, root.Right)

			if !isDepth {
				depth++
				isDepth = true
			}
		}
		i++

		isDepth = false
	}

	
	return Trees{
		Root:root,
		depth:depth,
		RWMutex:sync.RWMutex{},
	}
}

func (t *Trees) PreOrder() []int {
	res := make([]int, 0)
	toIntsPreOrder(t.Root, &res)
	return res
}

func (t *Trees) InOrder() []int {
	res := make([]int, 0)
	toIntsInOrder(t.Root, &res)
	return res
}

func (t *Trees) PostOrder() []int {
	res := make([]int, 0)
	toIntsPostOrder(t.Root, &res)
	return res
}

// 格式化tree root 以图像的形式展示出来
// 实质上还是中序遍历,只不过是输出字符串而已
func (t Trees) String() string {
	root := t.Root
	buf := bytes.NewBuffer([]byte("trees fmt \n"))

	treesString(root, buf, 0)
	return buf.String()
}

func treesString(root *TreeNode, buf *bytes.Buffer, depth int) {
	if root == nil {
		panic("cant fmt nil tree root")
	}

	if root.Left != nil {
		treesString(root.Left, buf, depth+1)
	}

	for i := 0; i < depth; i++ {
		buf.WriteString("     ")
	}
	buf.WriteString(fmt.Sprintf(" --- %d \n", root.Val))

	if root.Right != nil {
		treesString(root.Right, buf, depth+1)
	}
}

func (t *Trees) Depth() int {
	return t.depth
}

func (t *Trees) ptrDepth() int {
	return t.depth
}

func (t Trees) NormalDepth() int {
	return t.depth
}

// 前序遍历
func toIntsPreOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	*res = append(*res, root.Val)
	if root.Left != nil {
		toIntsPreOrder(root.Left, res)
	}

	if root.Right != nil {
		toIntsPreOrder(root.Right, res)
	}

}

// 迭代法前序遍历
func toIntsPreOrderIterate(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	node := root

	for node != nil || len(stack) != 0 {
		for node != nil {
			res = append(res, node.Val)		// Add before going to children
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node = node.Right
	}

	return res
}

// 中序遍历
func toIntsInOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	if root.Left != nil {
		toIntsInOrder(root.Left, res)
	}
	*res = append(*res, root.Val)

	if root.Right != nil {
		toIntsInOrder(root.Right, res)
	}

}

// 迭代法中序遍历
func toIntsInOrderIterate(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)

	stack := make([]*TreeNode, 0)
	node := root
	for node != nil || len(stack) != 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1]
		res = append(res, node.Val)		// Add after all left children
		stack = stack[:len(stack)-1]

		node = node.Right
	}

	return res
}

func toIntsPostOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	if root.Left != nil {
		toIntsPostOrder(root.Left, res)
	}

	if root.Right != nil {
		toIntsPostOrder(root.Right, res)
	}
	*res = append(*res, root.Val)
}

// 后续遍历,迭代方式
func toIntsPostOrderIterate(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)

	stack := make([]*TreeNode, 0)
	node := root
	for node != nil {
		for node != nil {
			res = append(res, node.Val)		// Add after all left children
			stack = append(stack, node)
			node = node.Right

		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node = node.Left
	}


	return res
}

