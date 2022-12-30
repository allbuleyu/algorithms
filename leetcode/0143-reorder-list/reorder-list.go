package prob0143

import "algorithms/kit"

type ListNode = kit.ListNode
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
	slow := head
	fast := head

	pre := &ListNode{Next:nil}
	for fast != nil && fast.Next != nil {
		// reverse the left-side nodes prepare to join the right-side nodes
		tmp := slow
		slow = slow.Next
		fast = fast.Next.Next

		tmp.Next = pre.Next
		pre.Next = tmp
	}

	head = pre.Next
	pre.Next = nil

	// list number is odd, and tail with the list is slow
	if fast != nil {
		tmp := slow
		slow = slow.Next

		tmp.Next = nil
		pre.Next = tmp
	}

	fast = head
	for slow != nil {
		tmpf := fast
		tmps := slow

		fast = fast.Next
		slow = slow.Next

		tmps.Next = pre.Next
		tmpf.Next = tmps
		pre.Next = tmpf
	}

}