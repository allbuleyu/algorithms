package challenge

import "github.com/allbuleyu/algorithms/kit"

// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/529/week-2/3290/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode = kit.ListNode
func middleNode(head *ListNode) *ListNode {
	var f, s *ListNode
	f, s = head, head

	for f != nil && f.Next != nil {
		f = f.Next.Next
		s = s.Next
	}

	return s
}
