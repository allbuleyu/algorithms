package prob0092

import "github.com/allbuleyu/algorithms/kit"

type ListNode = kit.ListNode
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}

	leftNode, rightNode := head, head

	var reverseHead *ListNode
	for n > 0 {
		if m > 0 {
			leftNode = leftNode.Next
			rightNode = rightNode.Next
		}else {
			tmp := rightNode
			rightNode = rightNode.Next

			reverseHead, tmp.Next = tmp, reverseHead
		}


		m--
		n--
	}

	leftNode.Next = reverseHead

	return head
}