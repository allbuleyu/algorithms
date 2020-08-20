package prob0146

type LRUCache struct {
	tail *myList
	mapping map[int]*myList
	capacity int
}

type myList struct {
	k, v int
	prev, next *myList
}

func Constructor(capacity int) LRUCache {
	return LRUCache{mapping: make(map[int]*myList), capacity: capacity}
}


func (this *LRUCache) Get(key int) int {
	if _, ok := this.mapping[key]; !ok {
		return -1
	}

	node := this.mapping[key]

	moveToHead(this.tail, node)
	this.tail = node

	return node.v
}

func moveToHead(tail, node *myList) {
	if tail == node {
		return
	}

	// if node.next == nil imply node is new NODE
	if node.next != nil {
		node.next.prev = node.prev
		node.prev.next = node.next
	}

	head := tail.next
	tail.next = node
	node.prev = tail

	head.prev = node
	node.next = head
}

func (this *LRUCache) Put(key int, value int)  {
	if node, ok := this.mapping[key]; ok {
		node.v = value

		moveToHead(this.tail, node)

		this.tail = node
		return
	}

	if this.capacity == 0 {
		next := this.tail.next
		delete(this.mapping, next.k)

		next.k = key
		next.v = value

		this.mapping[key] = next

		this.tail = next

		return
	}

	node := &myList{k: key, v: value}
	if this.tail == nil {
		node.prev = node
		node.next = node
		this.tail = node
	}else {
		moveToHead(this.tail, node)

		this.tail = node
	}

	this.capacity--
	this.mapping[key] = node
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */