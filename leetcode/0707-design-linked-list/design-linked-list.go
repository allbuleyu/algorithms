package prob0707

// 双链表的实现
type MyLinkedList struct {
	head *DoubleListNode
	l int
	tail *DoubleListNode
}

type DoubleListNode struct {
	Prev *DoubleListNode
	Val int
	Next *DoubleListNode
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}


/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.l {
		return -1
	}

	cur := this.head
	for index > 0 {
		cur = cur.Next
		index--
	}

	return cur.Val
}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
	head := &DoubleListNode{
		Prev: nil,
		Val: val,
		Next: this.head,
	}

	if this.head != nil {
		this.head.Prev = head
	}

	if this.l == 0 {
		this.tail = head
	}

	this.head = head
	this.l++
}


/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
	tail := &DoubleListNode{
		Prev: this.tail,
		Val:  val,
		Next: nil,
	}

	if this.tail != nil {
		this.tail.Next = tail
	}

	if this.l == 0 {
		this.head = tail
	}

	this.tail = tail
	this.l++
}


/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	if index == 0 {
		this.AddAtHead(val)
		return
	}

	if index == this.l {
		this.AddAtTail(val)
		return
	}

	next := this.head
	for index > 0 {
		next = next.Next
		index--
	}

	prev := next.Prev
	insert := &DoubleListNode{
		Prev: prev,
		Val:  val,
		Next: next,
	}

	prev.Next = insert
	next.Prev = insert
	this.l++
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
	if index < 0 || index >= this.l {
		return
	}

	cur := this.head
	if index == 0 {
		this.head = this.head.Next
	}

	if index == this.l-1 {
		this.tail = this.tail.Prev
	}

	for index > 0 {
		cur = cur.Next
		index--
	}

	prev := cur.Prev
	next := cur.Next

	if prev != nil {
		prev.Next = next
	}

	if next != nil {
		next.Prev = prev
	}

	this.l--
	if this.l == 0 {
		this.tail = nil
	}
}

// 单链表的的实现
//type MyLinkedList struct {
//	head *ListNode
//	l int
//	tail *ListNode
//}
//
//type ListNode struct {
//	Val int
//	Next *ListNode
//}
//
//
///** Initialize your data structure here. */
//func Constructor() MyLinkedList {
//	return MyLinkedList{}
//}
//
//
///** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
//func (this *MyLinkedList) Get(index int) int {
//	if index < 0 || index >= this.l {
//		return -1
//	}
//
//	cur := this.head
//	for index > 0 {
//		cur = cur.Next
//		index--
//	}
//
//	return cur.Val
//}
//
//
///** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
//func (this *MyLinkedList) AddAtHead(val int)  {
//	cur := &ListNode{
//		Val: val,
//		Next: this.head,
//	}
//
//	this.head = cur
//	if this.l == 0 {
//		this.tail = cur
//	}
//
//	this.l++
//}
//
//
///** Append a node of value val to the last element of the linked list. */
//func (this *MyLinkedList) AddAtTail(val int)  {
//	cur := &ListNode{
//		Val: val,
//		Next: nil,
//	}
//
//	if this.l == 0 {
//		this.head = cur
//		this.tail = cur
//		this.l++
//		return
//	}
//
//	this.tail.Next = cur
//	this.tail = this.tail.Next
//	this.l++
//}
//
//
///** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
//func (this *MyLinkedList) AddAtIndex(index int, val int)  {
//	if index < 0 || index > this.l {
//		return
//	}
//
//	if index == 0 {
//		this.AddAtHead(val)
//		return
//	}
//
//	if index == this.l {
//		this.AddAtTail(val)
//		return
//	}
//
//	pre := &ListNode{
//		Val:  0,
//		Next: this.head,
//	}
//
//	for index > 0 {
//		pre = pre.Next
//		index--
//	}
//
//	cur := &ListNode{
//		Val:  val,
//		Next: pre.Next,
//	}
//
//	pre.Next = cur
//	this.l++
//}
//
//
///** Delete the index-th node in the linked list, if the index is valid. */
//func (this *MyLinkedList) DeleteAtIndex(index int)  {
//	if index < 0 || index >= this.l {
//		return
//	}
//
//	pre := &ListNode{
//		Val:  0,
//		Next: this.head,
//	}
//
//	for i := index; i > 0; i-- {
//		pre = pre.Next
//	}
//
//	pre.Next = pre.Next.Next
//	this.l--
//
//	if index == 0 {
//		this.head = pre.Next
//	}
//
//	if this.l == index && this.l != 0 {
//		this.tail = pre
//	}
//
//	if this.l == 0 {
//		this.tail = nil
//	}
//}
//
//func (this *MyLinkedList) listToArr() []int {
//	head := this.head
//	nums := make([]int, 0)
//
//	for head != nil {
//
//		nums = append(nums, head.Val)
//		head = head.Next
//	}
//
//	return nums
//}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */