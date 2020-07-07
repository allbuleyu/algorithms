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
	helpRecursion(root)
	return root
}

// it is very easy to tranfer iteration
func helpRecursion(root *Node) {
	if root == nil {
		return
	}

	var child, prev *Node
	node := root
	for node != nil {
		if node.Left == nil && node.Right == nil {
			node = node.Next
			continue
		}

		if node.Left != nil {
			if child == nil {
				child = node.Left
				prev = node.Left
			}else {
				prev.Next = node.Left
				prev = prev.Next
			}
		}

		if node.Right != nil {
			if child == nil {
				child = node.Right
				prev = node.Right
			}else {
				prev.Next = node.Right
				prev = prev.Next
			}
		}

		node = node.Next
	}

	helpRecursion(child)
}
