package prob0061

import "github.com/allbuleyu/algorithms/kit"

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

	return iterate(head, k)
}

func iterate(head *ListNode, k int) *ListNode {
	lenList := 1
	pre := &ListNode{Next:head}
	cur := head

	for cur.Next != nil {
		lenList++
		cur = cur.Next
	}

	// return origin list
	if k % lenList == 0 {
		return head
	}

	// rotate list tail
	tail := cur

	// find the rotate position
	cur = pre
	for k = lenList-(k%lenList); k > 0; k-- {
		cur = cur.Next
	}

	head = cur.Next
	cur.Next = nil
	tail.Next=pre.Next
	pre.Next=head

	return pre.Next
}