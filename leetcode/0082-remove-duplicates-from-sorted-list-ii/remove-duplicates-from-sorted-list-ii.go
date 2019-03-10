package prob0082

import "github.com/allbuleyu/algorithms/kit"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = kit.ListNode
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next != nil {
		return  head
	}

	// 处理头结点重复的情况
	if  head.Val == head.Next.Val {
		for head.Next != nil && head.Val == head.Next.Val {
			head = head.Next
		}
		head = deleteDuplicates(head.Next)

		return head
	}

	// 头结点不重复的情况
	head.Next = deleteDuplicates(head.Next)

	return head
}