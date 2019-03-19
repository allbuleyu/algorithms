package prob0019

import "github.com/allbuleyu/algorithms/kit"

type ListNode = kit.ListNode
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	return solution2(head,n)
}

func solution2(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next:head}
	first, second := dummy, dummy

	for first != nil {
		if n < 0 {
			second = second.Next
		}

		first = first.Next
		n--
	}

	second.Next = second.Next.Next

	return dummy.Next
}

func solution1(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	fast := head
	pre := &ListNode{Next:head}
	l := 0

	for fast != nil && fast.Next != nil {
		l++

		fast = fast.Next.Next
	}

	if fast == nil {
		l *= 2
	}else {
		l = l * 2 + 1
	}

	cur := pre
	for l-n>=0 {
		if l-n == 0 {
			cur.Next = cur.Next.Next
			break
		}

		l--
		cur = cur.Next
	}


	return pre.Next
}