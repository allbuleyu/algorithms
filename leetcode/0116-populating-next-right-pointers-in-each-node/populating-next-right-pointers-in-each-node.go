package prob0116

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

func helpRecursion(root *Node) {
	if root == nil || root.Left == nil {
		return
	}

	node := root
	for node != nil {
		node.Left.Next = node.Right
		if node.Next != nil {
			node.Right.Next = node.Next.Left
		}
		node = node.Next
	}

	helpRecursion(root.Left)
}

func helpTwoPointers(root *Node) {
	if root == nil || root.Left == nil {
		return
	}

	child := root
	node := root
	//for node != nil && node.Left != nil {
	//	node.Left.Next = node.Right
	//	if node.Next != nil {
	//		node.Right.Next = node.Next.Left
	//	}
	//
	//	node = node.Next
	//
	//	// 如果一层已经遍历完了,那就到下一层
	//	if node == nil {
	//		node = child.Left
	//		child = child.Left
	//	}
	//}

	// 这样可读性更好
	for child.Left != nil {
		for node != nil {
			node.Left.Next = node.Right
			if node.Next != nil {
				node.Right.Next = node.Next.Left
			}
			node = node.Next
		}
		child = child.Left
	}
}

