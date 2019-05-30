package prob0082

import "github.com/allbuleyu/algorithms/kit"

type ListNode = kit.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	return iterate(head)
}

func iterate(head *ListNode) *ListNode {
	cur := head
	pre := &ListNode{Next:head}
	dummy := pre

	for cur != nil {
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		}

		if pre.Next == cur {
			pre = pre.Next
		}else {
			pre.Next = cur.Next
		}

		cur = cur.Next
	}

	return dummy.Next
}

func recursion(head *ListNode) *ListNode {
	cur := head
	pre :=&ListNode{Next:head}
	dummy := pre

	for cur != nil {
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		}

		if pre.Next == cur {
			pre = pre.Next
		}else {
			pre.Next = recursion(cur.Next)
		}

		cur = cur.Next
	}

	return dummy.Next
}

