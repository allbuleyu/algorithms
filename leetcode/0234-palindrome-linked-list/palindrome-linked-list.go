package prob0234

import (
	"algorithms/kit"
)

type ListNode = kit.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	return solution1(head)
}

func reverse(h *ListNode) *ListNode {
	var reverseNode *ListNode

	for h != nil {
		tmp := h
		h = h.Next

		reverseNode, tmp.Next = tmp, reverseNode
	}

	return reverseNode
}

// 此解决方案速度100%,空间20%
// 尝试其它方案能不能100%速度,100%空间
func solution1(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow := head
	fast := head

	var converseNode, tmp *ListNode

	for fast != nil && fast.Next != nil {
		tmp = slow

		slow = slow.Next
		fast = fast.Next.Next

		tmp.Next = converseNode		// 这一步会导致head.next直接改变,所以要在这之前,让slow,fast走到下个节点
		converseNode = tmp
	}

	// 判断链表的长度是否为奇数,奇数要略过中间的数
	if fast != nil && fast.Next == nil {
		slow = slow.Next
	}

	for slow != nil {
		if slow.Val != converseNode.Val {
			return false
		}

		slow = slow.Next
		converseNode = converseNode.Next
	}

	return true
}