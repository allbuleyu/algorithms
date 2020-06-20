package prob0430

type Node struct {
	Val int
	Prev, Next, Child *Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */

func flatten(root *Node) *Node {
	return helpRecursion(root)
}

func helpRecursion(root *Node) *Node {
	if root == nil {
		return nil
	}

	//     if root.Child == nil {
	//         root.Next = helpRecursion(root.Next)
	//         return root
	//     }

	//     next := root.Next
	//     child := root.Child
	//     root.Child = nil
	//     root.Next = child
	//     child.Prev = root
	//     for child != nil {
	//         if child.Child != nil {
	//             child = helpRecursion(child)
	//         }

	//         if child.Next == nil {
	//             break
	//         }
	//         child = child.Next
	//     }
	//     child.Next = next
	//     if next != nil {
	//         next.Prev = child
	//     }

	// 这种解法更加的直观
	cur := root
	for cur != nil {
		if cur.Child != nil {
			next := cur.Next
			child := cur.Child
			cur.Child = nil

			cur.Next = child
			child.Prev = cur

			child = helpRecursion(child)

			last := child
			for last.Next != nil {
				last = last.Next
			}

			last.Next = next
			if next != nil {
				next.Prev = last
			}

			cur = last
		}

		cur = cur.Next
	}

	return root
}

