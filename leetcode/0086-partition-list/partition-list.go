package prob0086

import "algorithms/kit"

type ListNode = kit.ListNode
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 还有一种不好理解的方式,交换节点法,效率一样
func partition(head *ListNode, x int) *ListNode {
	var less, lessHead,great, greatHead *ListNode
	for head != nil {
		if head.Val < x {
			if lessHead == nil {
				lessHead = head
				less = head
			}else {
				less.Next = head
				less = less.Next

			}
		}else {
			if greatHead == nil {
				greatHead = head
				great = head
			}else {
				great.Next = head
				great = great.Next

			}
		}

		head = head.Next
	}

	if great != nil {
		great.Next=nil
	}

	if lessHead == nil {
		return greatHead
	}

	less.Next = greatHead
	return lessHead
}