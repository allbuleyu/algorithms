package prob0138

type Node struct {
	Val int
	Next, Random *Node
}
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	return helpFunc2(head)
}

// We can avoid using extra space for old node ---> new node mapping, by tweaking the original linked list.
// Simply interweave the nodes of the old and copied list. For e.g.
// Old List: A --> B --> C --> D
// InterWeaved List: A --> A' --> B --> B' --> C --> C' --> D --> D'
func helpFunc2(head *Node) *Node {
	if head == nil {
		return nil
	}

	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = &Node{Val:cur.Val, Next:next}

		cur = next
	}

	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}

		cur = cur.Next.Next
	}

	cur = head
	cpHead := &Node{}
	cpCur := cpHead
	for cur != nil {
		cpCur.Next = cur.Next
		cur.Next = cur.Next.Next

		cur = cur.Next
		cpCur = cpCur.Next
	}

	return cpHead.Next
}

func helpFunc1(head *Node) *Node {
	if head == nil {
		return nil
	}

	h := make(map[*Node]int)
	ch := make(map[int]*Node)
	cpHead := &Node{Val:head.Val}
	cpCur := cpHead
	cur := head.Next

	h[cur] = 0
	ch[0] = cpCur
	i := 1

	for cur != nil {
		cpCur.Next = &Node{Val:cur.Val}
		h[cur] = i
		ch[i] = cpCur.Next
		i++

		cur = cur.Next
		cpCur = cpCur.Next
	}

	cur = head
	cpCur = cpHead
	for cur != nil {
		if cur.Random != nil {
			cpCur.Random = ch[h[cur.Random]]
		}
		cur = cur.Next
		cpCur = cpCur.Next
	}

	return cpHead
}