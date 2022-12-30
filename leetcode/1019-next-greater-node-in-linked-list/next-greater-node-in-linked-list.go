package prob1019

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
func nextLargerNodes(head *ListNode) []int {
	if head == nil {
		return nil
	}

	return iterateHelp(head)
}

func iterateHelp(head *ListNode) []int {
	var cur *ListNode
	ans := make([]int, 0)

	for head != nil {
		cur = head
		for cur != nil {
			// iterate to the tail, not find next larger node
			if cur.Next == nil {
				ans = append(ans, 0)
				break
			}

			// find the next larger node, append to ans, and break to the loop
			if cur.Next.Val > head.Val {
				ans = append(ans, cur.Next.Val)
				break
			}

			cur = cur.Next
		}

		head = head.Next
	}

	return ans
}

func recursionHelp(head *ListNode) []int {
	if head == nil {
		return []int{}
	}

	ans := make([]int, 0)

	cur := head
	for cur != nil {
		if cur.Next == nil {
			ans = append(ans, 0)
			break
		}

		if cur.Next.Val > head.Val {
			ans = append(ans, cur.Next.Val)
			break
		}

		cur = cur.Next
	}

	ans = append(ans, recursionHelp(head.Next)...)

	return ans
}
