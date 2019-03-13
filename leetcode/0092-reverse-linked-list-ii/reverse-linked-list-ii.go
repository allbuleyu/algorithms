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
	preHead := &ListNode{Next:head}
	leftNode, rightNode := preHead, preHead.Next

	var reverseHead, reverseTail *ListNode
	for n > 0 {
		if m > 1 {
			leftNode = leftNode.Next
			rightNode = rightNode.Next
		}else {
			tmp := rightNode
			rightNode = rightNode.Next

			reverseHead, tmp.Next = tmp, reverseHead
			if reverseTail == nil {
				reverseTail = reverseHead
			}
		}

		m--
		n--
	}

	leftNode.Next = reverseHead
	reverseTail.Next = rightNode

	return preHead.Next
}