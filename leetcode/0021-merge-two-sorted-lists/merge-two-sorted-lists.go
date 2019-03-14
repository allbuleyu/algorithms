package prob0021

import "github.com/allbuleyu/algorithms/kit"

type ListNode = kit.ListNode
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 有一条链为nil，直接返回另一条链
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	preHead := &ListNode{}
	cur := preHead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
			cur = cur.Next
		}else {
			cur.Next = l2
			l2 = l2.Next
			cur = cur.Next
		}
	}

	if l2 != nil {
		l1 = l2
	}

	for l1 != nil {
		cur.Next = l1
		l1 = l1.Next
		cur = cur.Next
	}

	return preHead.Next
}