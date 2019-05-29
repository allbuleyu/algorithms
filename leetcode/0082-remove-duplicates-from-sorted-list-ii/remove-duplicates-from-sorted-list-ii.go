package prob0082

import "github.com/allbuleyu/algorithms/kit"

type ListNode = kit.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	return iterate(head)
}

func iterate(head *ListNode) *ListNode {

}

func recursion(head *ListNode) *ListNode {

}

