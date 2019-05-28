package prob0160

import "github.com/allbuleyu/algorithms/kit"

type ListNode = kit.ListNode
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pa := headA
	pb := headB
	travelNum := 0
	for pa != nil || pb != nil {
		if pa == pb {
			return pa
		}

		if pa == nil {
			pa = headB
			travelNum++
		}

		if travelNum == 2 {
			break
		}

		if pb == nil {
			pb = headA
		}

		pa = pa.Next
		pb = pb.Next
	}

	return nil
}