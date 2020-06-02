package kit

type priorityQueue struct {
	queue []*TreeNode
}

func (p *priorityQueue) Put(node *TreeNode) {

	panic("implement me")
}

func (p *priorityQueue) Pop() *TreeNode {
	panic("implement me")
}

func (p *priorityQueue) Down(root *TreeNode) *TreeNode {
	panic("implement me")
}

func (p *priorityQueue) Up(node *TreeNode) *TreeNode {
	panic("implement me")
}

type Priorityer interface {
	Put(node *TreeNode)
	Pop() *TreeNode
	Down(root *TreeNode) *TreeNode
	Up(node *TreeNode) *TreeNode
}

func NewPriorityQueue(nums []int) *priorityQueue {
	return nil
}




