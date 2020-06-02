package prob0206

import "github.com/allbuleyu/algorithms/kit"

type ListNode = kit.ListNode
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	dummy := &ListNode{Next:nil}

	for head != nil {
		tmp := head.Next
		head.Next = dummy.Next
		dummy.Next = head
		head = tmp
	}

	return dummy.Next
}