package prob0024

import "algorithms/kit"

type ListNode = kit.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	return recursion(head)
}

func iterate(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	num := 0
	cur := head
	opCur := &ListNode{Next:head}
	head = head.Next

	for cur != nil {
		num++
		cur = cur.Next

		if num % 2 == 0 && num > 0 {
			tmp := opCur.Next.Next
			opCur.Next.Next = tmp.Next
			tmp.Next = opCur.Next
			opCur.Next = tmp

			opCur = opCur.Next.Next
		}
	}

	return head
}

func recursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	n := head.Next
	head.Next = recursion(head.Next.Next)
	n.Next = head

	return n
}