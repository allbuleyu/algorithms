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
	return recursion1(head, k)
}

func iterate(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	num := 0
	cur := head
	pre := &ListNode{Next: head}

	for cur != nil {
		num++
		cur = cur.Next

		if num%k == 0 {
			groupHead := pre.Next
			newGroupHead := reverseNode(pre.Next, cur)
			pre.Next = newGroupHead

			if num == k {
				head = newGroupHead
			}

			groupHead.Next = cur
			pre = groupHead
		}
	}

	return head
}

func reverseNode(head *ListNode, end *ListNode) *ListNode {
	pre := &ListNode{Next: nil}
	for head != nil && head != end {
		tmp := head
		head = head.Next

		tmp.Next = pre.Next
		pre.Next = tmp
	}

	return pre.Next
}

// 使用的内存比迭代的多了50%
func recursion(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}

	num := 0
	cur := head

	for cur != nil && num != k {
		cur = cur.Next
		num++
	}

	if num == k {
		newGroupHead := reverseNode(head, cur)

		head.Next = recursion(cur, k)
		head = newGroupHead
	}

	return head
}

func recursion1(head *ListNode, k int) *ListNode {
	if head == nil || k < 2 {
		return head
	}

	cur := head
	i := 0
	for i < k && cur != nil {
		cur = cur.Next
		i++
	}

	if i < k {
		return head
	}

	var rHead *ListNode
	cur = head
	for i = 0; i < k; i++ {
		next := cur.Next
		cur.Next = rHead
		rHead = cur

		cur = next
	}
	head.Next = recursion1(cur, k)

	return rHead
}
