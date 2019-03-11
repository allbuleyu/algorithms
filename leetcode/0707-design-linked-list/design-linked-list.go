package prob0707


type MyLinkedList struct {
	head *Node
}

type Node struct {
	val interface{}
	next *Node
}


/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{head:&Node{val:0}}
}


/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	cur := this.head
	for cur != nil && index >= 0 {
		if index == 0 {
			return cur.val.(int)
		}

		cur = cur.next
		index--
	}

	return -1
}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
	head := this.head
	cur := &Node{val:val, next:head}

	this.head = cur
}


/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
	cur := this.head
	for cur != nil {
		if cur.next == nil {
			break
		}

		cur = cur.next
	}

	cur.next = &Node{val:val}
}


/** Add a node of value val before the index-th node in the linked list.
 *If index equals to the length of linked list, the node will be appended to the end of linked list.
 *If index is greater than the length, the node will not be inserted.
 */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	if index < 0 {
		return
	}
	cur := this.head

	if index == 0 {
		this.head = &Node{val:val, next:cur}
	}

	headPrev := &Node{val:nil, next:this.head}

	for headPrev != nil && index >= 0 {
		if index == 0 {
			break
		}

		headPrev = headPrev.next
		index--
	}

	if headPrev == nil {
		return
	}

	headPrev.next = &Node{val:val, next:headPrev.next}
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
	if index < 0 {
		return
	}

	if index == 0 {
		this.head = this.head.next
	}

	headPrev := &Node{next:this.head}
	for index >= 0 && headPrev.next != nil {
		if index == 0 {
			headPrev.next = headPrev.next.next
		}

		headPrev = headPrev.next
		index--
	}
}

func (this *MyLinkedList) FmtList() {

}

func (this *MyLinkedList) ConvertToArr() []int {
	cur := this.head
	res := make([]int, 0)

	for cur != nil {
		v, ok := cur.val.(int)
		if ok {
			res = append(res, v)
		}
		cur = cur.next
	}

	return res
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */