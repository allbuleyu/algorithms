package kit

type ListNode struct {
	Val int
	Next *ListNode
}

func (l *ListNode) Remove(v int) {
	
}

func CreateNodes(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := new(ListNode)

	tmp := head
	for i := 0; i < len(nums); i++ {
		tmp.Next = &ListNode{Val:nums[i]}
		tmp = tmp.Next
	}
	
	return head.Next
}

func CreateNodesWithPosCircle(nums []int, pos int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := new(ListNode)

	var circleNode *ListNode
	tmp := head
	for i := 0; i < len(nums); i++ {
		tmp.Next = &ListNode{Val:nums[i]}
		tmp = tmp.Next

		if pos == 0 {
			circleNode = tmp
		}
		pos--
	}

	tmp.Next = circleNode

	return head.Next
}

func TransferNodes(head *ListNode) []int {
	nums := make([]int, 0)

	for head != nil {

		nums = append(nums, head.Val)
		head = head.Next
	}

	return nums
}