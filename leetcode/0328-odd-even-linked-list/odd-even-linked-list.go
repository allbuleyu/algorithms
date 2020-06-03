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


func helpStupid(head *ListNode) *ListNode {
	var oddTail, evenHead *ListNode

	oddTail = &ListNode{Next:nil}
	evenHead = nil
	evenTail := &ListNode{Next:nil}

	cur := head
	for i := 1; cur != nil; i++ {
		if i % 2 != 0 {
			oddTail.Next = cur
			oddTail = oddTail.Next

		}else {
			if evenHead == nil {
				evenHead = cur
			}

			evenTail.Next = cur
			evenTail = evenTail.Next
		}
		cur = cur.Next
	}

	if evenTail != nil {
		evenTail.Next = nil
	}

	oddTail.Next = evenHead

	return head
}