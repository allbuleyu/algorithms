package prob1171

import "algorithms/kit"

type ListNode = kit.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeZeroSumSublists(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	cur := dummy
	prevSum := 0
	seen := make(map[int]*ListNode)
	for cur != nil {
		prevSum += cur.Val
		seen[prevSum] = cur
		cur = cur.Next
	}

	prevSum = 0
	cur = dummy
	for cur != nil {
		prevSum += cur.Val
		cur.Next = seen[prevSum].Next
		cur = cur.Next
	}

	return dummy.Next
}

// todo:没完成
func removeZeroSumSubArray(nums []int) []int {
	seen := make(map[int]int)
	seen[0] = -1
	prevSum := 0
	for i := 0; i < len(nums); i++ {
		prevSum = nums[i] + prevSum
	}

	res := make([]int, 0)
	prevSum = 0
	for i := 0; i < len(nums); i++ {
		prevSum += nums[i]
		j := seen[prevSum]
		if j >= 0 {
			res = append(res, nums[i:j+1]...)
		}
	}

	return res
}
