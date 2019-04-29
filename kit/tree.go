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
	Head *TreeNode
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
		node := queue[0]
		queue = queue[1:]

		if ints[i] != Null {
			node.Left = &TreeNode{Val:ints[i]}
			queue = append(queue, node.Left)

			if !isDepth {
				depth++
				isDepth = true
			}
		}
		i++

		if i < n && ints[i] != Null {
			node.Right = &TreeNode{Val:ints[i]}
			queue = append(queue, node.Right)

			if !isDepth {
				depth++
				isDepth = true
			}
		}
		i++

		isDepth = false
	}

	
	return Trees{
		Head:root,
		depth:depth,
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

// 显示出来的图形是前序遍历的结果
func (t *Trees) String() string {
	root := t.Head
	stack := make([]*TreeNode, 0)
	depth := t.depth

	buf := bytes.NewBuffer([]byte("tree fmt \n"))

	fmtStr := "-------"
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		cur = stack[len(stack)-1]

		stack = stack[:len(stack)-1]
		buf.WriteString(fmt.Sprintf("%s ;%d; %d \n", fmtStr, depth, cur.Val))
		cur = cur.Right
	}


	return buf.String()
}

func (t *Trees) Depth() int {
	return t.depth
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

