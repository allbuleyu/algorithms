package prob0083

import "github.com/allbuleyu/algorithms/kit"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode = kit.ListNode
func deleteDuplicates(head *ListNode) *ListNode {
	op := head

	noDup := head.Val
	for op != nil {
		if op.Next == nil {
			break
		}

		if op.Next.Val == noDup {
			op.Next = op.Next.Next
		}else {
			noDup = op.Next.Val
		}

		op = op.Next
	}

	return head
}