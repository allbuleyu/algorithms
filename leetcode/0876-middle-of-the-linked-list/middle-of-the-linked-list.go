package prob0876

import "github.com/allbuleyu/algorithms/kit"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode = kit.ListNode
func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur := head
	n := 0
	for cur != nil {
		cur = cur.Next
		n++
	}

	n = n/2
	for n > 0 {
		head = head.Next
		n--
	}

	return head
}