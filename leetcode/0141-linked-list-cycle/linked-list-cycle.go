package prob0141

import "algorithms/kit"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = kit.ListNode
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	var slow, fast = head, head.Next

	//for slow != nil && (fast != nil && fast.Next != nil) {
	//	if slow == fast {
	//		return true
	//	}
	//
	//	slow = slow.Next
	//	fast = fast.Next.Next
	//}

	// 显然,此判断更优雅,更精确
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}