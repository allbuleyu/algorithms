package prob0142

import "algorithms/kit"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = kit.ListNode


// http://www.cnblogs.com/hiddenfox/p/3408931.html 单向链表的各种变形
func detectCycle(head *ListNode) *ListNode {
	return helpFun1(head)
}

func helpFun1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	s, f := head, head
	for {
		if f == nil || f.Next == nil {
			return nil
		}

		s = s.Next
		f = f.Next.Next

		if f == s {
			break
		}
	}

	s = head
	for s != f {
		s = s.Next
		f = f.Next
	}

	return s
}

func helpFunc(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	var slow, fast = head, head
	for {
		if fast.Next == nil || fast.Next.Next == nil {
			return nil
		}

		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			break
		}
	}

	//Input:	此测试用例没过
	//[1,2]
	//0
	//slow = head
	//
	//for {
	//	slow = slow.Next
	//	fast = fast.Next
	//
	//	if slow == fast {
	//		break
	//	}
	//}

	slow = head

	for {
		if slow == fast {
			break
		}

		slow = slow.Next
		fast = fast.Next
	}

	return fast
}