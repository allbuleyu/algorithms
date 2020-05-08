package prob0146

//Design and implement a data structure for Least Recently Used (LRU) cache. It should support the following operations: get and put.
//
//get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
//put(key, value) - Set or insert the value if the key is not already present. When the cache reached its capacity, it should invalidate the least recently used item before inserting a new item.
//
//The cache is initialized with a positive capacity.
//
//Follow up:
//Could you do both operations in O(1) time complexity?
//
//Example:
//
//LRUCache cache = new LRUCache( 2 /* capacity */ );
//
//cache.put(1, 1);
//cache.put(2, 2);
//cache.get(1);       // returns 1
//cache.put(3, 3);    // evicts key 2
//cache.get(2);       // returns -1 (not found)
//cache.put(4, 4);    // evicts key 1
//cache.get(1);       // returns -1 (not found)
//cache.get(3);       // returns 3
//cache.get(4);       // returns 4


type LRUCache struct {
	head *node
	tail *node
	store map[int]*node
	cap int
}

type node struct {
	prev *node
	next *node
	key int
	val int
}


func Constructor(capacity int) LRUCache {
	return LRUCache{
		store: make(map[int]*node),
		cap:   capacity,
	}
}


func (this *LRUCache) Get(key int) int {
	store := this.store
	if nd, ok := store[key];ok {


		return nd.val
	}

	return -1
}


func (this *LRUCache) Put(key int, value int)  {

}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */