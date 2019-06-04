package prob0725

import (
	"github.com/allbuleyu/algorithms/kit"
)

type ListNode = kit.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitListToParts(root *ListNode, k int) []*ListNode {

	return help(root, k)
}

func help(root *ListNode, k int) []*ListNode {
	n := 0
	res := make([]*ListNode, k)
	cur := root

	for cur != nil {
		n++
		cur = cur.Next
	}

	width := n/k
	Remainder := n%k

	cur = root
	for i := 0; i < k; i++ {
		head := cur

		// that's important!
		var xx int = 0
		if i < Remainder {
			xx = 1
		}
		for j := 0; j < width+xx-1; j++ {
			if cur != nil {
				cur = cur.Next
			}
		}

		if cur != nil {
			tmp := cur
			cur = cur.Next
			tmp.Next = nil
		}

		res[i] = head
	}

	return res
}