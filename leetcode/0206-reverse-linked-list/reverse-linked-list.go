package prob0206

import "algorithms/kit"

type ListNode = kit.ListNode
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	dummy := &ListNode{}
	helpRecursionDownUp(head, dummy)
	return dummy.Next
}

func helpIteration(head *ListNode) *ListNode {
	dummy := &ListNode{Next:nil}

	for head != nil {
		tmp := head.Next
		head.Next = dummy.Next
		dummy.Next = head
		head = tmp
	}

	return dummy.Next
}

func helpRecursionDownUp(head, newHead *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		newHead.Next = head
		return head
	}

	next := helpRecursionDownUp(head.Next, newHead)
	next.Next = head
	head.Next = nil
	return head
}

func helpRecursionUpDown(head, dummy *ListNode) *ListNode {
	if head == nil {
		return head
	}

	next := head.Next
	head.Next = dummy.Next
	dummy.Next = head

	return helpRecursionUpDown(next, dummy)
}