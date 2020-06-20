package prob0117

type Node struct {
	Val int
	Left,Right,Next *Node
}
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	helpTwoPointers(root)
	return root
}

func helpTwoPointers(root *Node) {
	var start, cur *Node
	start = root
	for start != nil {
		cur = start
		for cur != nil && (cur.Left != nil || cur.Right != nil || cur.Next != nil) {
			var pre *Node
			if cur.Left != nil {
				pre = cur.Left
				if cur.Right != nil {
					cur.Left.Next = cur.Right
					pre = cur.Right
				}

			}else if cur.Right != nil {
				pre = cur.Right
			}else {
				cur = cur.Next
				continue
			}

			if cur.Next != nil {
				next := cur.Next
				for next != nil {
					if next.Left != nil {
						pre.Next = next.Left
						break
					}else if next.Right != nil {
						pre.Next = next.Right
						break
					}

					next = next.Next
				}
			}

			cur = cur.Next
		}

		if start.Left != nil {
			start = start.Left
		}else if start.Right != nil {
			start = start.Right
		}else {
			start = start.Next
		}
	}
}