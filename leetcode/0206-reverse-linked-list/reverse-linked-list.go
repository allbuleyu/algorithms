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
	if head == nil || head.Next == nil {
		return head
	}

	var reverseHead *ListNode
	for head != nil {

		tmp := head
		head = head.Next

		tmp.Next = reverseHead
		reverseHead = tmp
		//reverseHead, tmp.Next = tmp, reverseHead
	}

	return reverseHead
}