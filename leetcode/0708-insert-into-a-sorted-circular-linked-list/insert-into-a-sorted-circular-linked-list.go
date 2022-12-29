package prob0708

import "algorithms/kit"

type Node = kit.ListNode
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 * }
 */

func insert(aNode *Node, x int) *Node {
	if aNode == nil {
		head := &Node{Val:x}
		head.Next = head
		return head
	}

	min := aNode
	max := aNode
	cur := aNode
	for {
		if cur.Val < min.Val {
			min = cur
		}

		if cur.Val >= max.Val {
			max = cur
		}

		if cur.Val <= x && cur.Next.Val > x {
			next := cur.Next
			cur.Next = &Node{Val:x, Next:next}
			return aNode
		}

		if cur.Next == aNode {
			break
		}
		cur = cur.Next
	}

	next := max.Next
	max.Next = &Node{Val:x, Next:next}

	return aNode
}