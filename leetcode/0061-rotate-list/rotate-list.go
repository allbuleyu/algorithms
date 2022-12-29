package prob0061

import "algorithms/kit"

type ListNode = kit.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	return optimize(head, k)
}

func optimize(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	l := 1
	cur := head
	for cur != nil && cur.Next != nil {
		l++
		cur = cur.Next
	}

	i := l - k % l
	if i == l {
		return head
	}

	cur.Next = head
	for ; i > 0; i-- {
		cur = cur.Next
	}

	head = cur.Next
	cur.Next = nil
	return head
}

func iterate(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	l := 0
	first := head
	for first != nil {
		l++
		first = first.Next
	}

	first = head
	second := head
	i := l - k % l
	if i == l {
		return head
	}

	for ; i > 1; i-- {
		first = first.Next
	}

	head = first.Next
	// first.Next 是新的头部,那么first就是新的尾部
	first, first.Next = first.Next, nil
	for first != nil && first.Next != nil {
		first = first.Next
	}

	first.Next = second

	return head
}