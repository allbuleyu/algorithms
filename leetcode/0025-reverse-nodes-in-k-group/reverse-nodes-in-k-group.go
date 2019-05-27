package prob0025

import "github.com/allbuleyu/algorithms/kit"

type ListNode = kit.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	return iterate(head, k)
}

func iterate(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	num := 0
	cur := head
	pre := &ListNode{Next:head}

	for cur != nil {
		num++
		cur = cur.Next

		if num % k == 0 {
			groupHead := pre.Next
			pre.Next = reverseNode(pre.Next)

			if num == k {
				head = pre.Next
			}

			groupHead.Next = cur
			pre.Next = cur
		}
	}

	return head
}

func reverseNode(head *ListNode) *ListNode {
	pre := &ListNode{Next:head}
	for head != nil {
		tmp := head
		head = head.Next

		tmp.Next = pre.Next
		pre.Next = tmp

	}

	return pre.Next
}