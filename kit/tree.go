package kit

type TreeInts []int

const Null int = -1 << 32

func (TreeInts) String() string {
	panic("implement me")
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}




