package prob0380

import "math/rand"

type RandomizedSet struct {
	hs map[int]int
	data []int
	curIndex int
}


/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		hs : make(map[int]int),
		data : make([]int, 0),
	}
}


/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.hs[val]; ok {
		return false
	}

	if this.curIndex == len(this.data) {
		this.hs[val] = this.curIndex
		this.data = append(this.data, val)
		this.curIndex++
		return true
	}

	this.hs[val] = this.curIndex
	this.data[this.curIndex] = val
	this.curIndex++

	return true
}


/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.hs[val]; !ok {
		return false
	}

	//for i := this.hs[val]+1; i < len(this.data); i++ {
	//	this.data[i-1] = this.data[i]
	//	this.hs[this.data[i]] = i-1
	//}

	// 互换当前要删除元素与最后元素的位置,并且把数组长度减一就可以
	index := this.hs[val]
	this.data[index] = this.data[this.curIndex-1]
	this.hs[this.data[index]] = index

	delete(this.hs, val)
	this.curIndex--

	return true
}


/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	if this.curIndex == 1 {
		return this.data[0]
	}

	return this.data[rand.Int() % this.curIndex]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */