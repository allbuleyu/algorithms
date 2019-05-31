package prob0328


import "github.com/allbuleyu/algorithms/kit"

type ListNode = kit.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	return iterateOptimize(head)
}

func iterate(head *ListNode) *ListNode {
	odds := &ListNode{}
	even := &ListNode{}
	evenHead := head.Next

	cur := head
	k := 1

	for cur != nil {
		if k % 2 == 0 {
			even.Next = cur
			even = even.Next

		}else {
			odds.Next = cur
			odds = odds.Next
		}

		k++
		cur = cur.Next
	}
	odds.Next = evenHead
	even.Next = nil

	return head
}

func iterateOptimize(head *ListNode) *ListNode {
	odds := head
	even := head.Next
	eHead := head.Next

	for even != nil && even.Next != nil {
		odds.Next = even.Next
		odds = odds.Next

		even.Next = odds.Next
		even = even.Next
	}
	odds.Next = eHead

	return head
}