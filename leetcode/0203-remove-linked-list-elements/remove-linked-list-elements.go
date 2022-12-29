package prob0203

import "algorithms/kit"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode = kit.ListNode

func removeElements(head *ListNode, val int) *ListNode {
	headPre := ListNode{Next: head}

	temp := &headPre
	for temp.Next != nil {
		if temp.Next.Val == val {
			// 删除符合条件的节点
			temp.Next = temp.Next.Next
		} else {
			temp = temp.Next
		}
	}

	return headPre.Next
}