package prob0019

import "algorithms/kit"

type ListNode = kit.ListNode
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	return twoPointers(head, n)
}

// len list = n + first->nil 的长度
func twoPointers(head *ListNode, n int) *ListNode {
	pre := &ListNode{
		Next:head,
	}

	first, second := pre, pre
	for i := 0; i < n+1; i++ {
		first = first.Next
	}

	for first != nil {
		first = first.Next
		second = second.Next
	}

	second.Next = second.Next.Next
	return pre.Next
}

func helpFunc1(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	pre := &ListNode{
		Next:head,
	}
	s, f := pre, head
	l := 0
	for f != nil && f.Next != nil {
		l++
		f = f.Next.Next
	}

	// f.Next == nil 表示是奇数长度
	totalLen := l * 2 + 1

	// f == nil 表示偶数长度
	if f == nil {
		totalLen--
	}

	for i := totalLen-n; i > 0; i-- {
		s = s.Next
	}

	// del
	s.Next = s.Next.Next

	return pre.Next
}