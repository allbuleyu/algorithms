package prob0234

import (
	"github.com/allbuleyu/algorithms/kit"
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
	var slow, fast *ListNode
	slow = head
	fast = head

	var converseNode *ListNode

	for fast != nil && fast.Next != nil {
		tmp := slow
		slow = slow.Next
		fast = fast.Next.Next

		tmp.Next = converseNode
		converseNode = tmp

		//fast = fast.Next.Next
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