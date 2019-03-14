package kit

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func Int2BSTree(nums []int) *TreeNode {
	return nil
}

func Ints2Tree(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}

	root := &TreeNode{Val:nums[0]}
	queue := make([]*TreeNode, 1, n)
	queue[0] = root

	for i := 1; i < n; {
		curNode := queue[0]
		queue = queue[1:]

		curNode.Left = &TreeNode{Val:nums[i]}
		queue = append(queue, curNode.Left)
		i++

		if i < n {
			curNode.Right = &TreeNode{Val:nums[i]}
			queue = append(queue, curNode.Right)
		}

		i++
	}

	return root
}

// 先序遍历法,把树转换成数组
func (root *TreeNode) PreTree2Int(nums []int) {
	nums = append(nums, root.Val)
	root.Left.PreTree2Int(nums)
	root.Right.PreTree2Int(nums)
}
