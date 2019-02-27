package prob0203

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	r := head
	for head.Next != nil {
		if head.Next.Val == val {
			head.Next = head.Next.Next
		}
		head = head.Next
	}

	if r.Val == val {
		r = r.Next
	}

	return r
}