package challenge

//https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/529/week-2/3292/

//type MinStack struct {
//	Value int
//	Next *MinStack
//	Min int
//}
//
//
///** initialize your data structure here. */
//func Constructor() MinStack {
//	return MinStack{}
//}
//
//
//func (this *MinStack) Push(x int)  {
//	next := this.Next
//	min := 0
//	if this.Next == nil {
//		min = x
//	}else {
//		if x < this.Next.Min {
//			min = x
//		}else {
//			min = this.Next.Min
//		}
//	}
//
//	this.Next = &MinStack{
//		Value: x,
//		Next:  next,
//		Min:   min,
//	}
//}
//
//
//func (this *MinStack) Pop()  {
//	if this.Next == nil {
//		return
//	}
//
//	this.Next = this.Next.Next
//}
//
//
//func (this *MinStack) Top() int {
//
//	return this.Next.Value
//}
//
//
//func (this *MinStack) GetMin() int {
//	return this.Next.Min
//}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */


// 数组方式实现
type MinStack struct {
	data []int
	min []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data: make([]int, 0),
		min:  make([]int, 0),
	}
}


func (this *MinStack) Push(x int)  {
	min := x
	if len(this.min) > 0 && x > this.min[len(this.min)-1] {
		min = this.min[len(this.min)-1]
	}

	this.data = append(this.data, x)
	this.min = append(this.min, min)
}


func (this *MinStack) Pop()  {
	if len(this.data) == 0 {
		return
	}
	this.data = this.data[:len(this.data)-1]
	this.min = this.min[:len(this.min)-1]
}


func (this *MinStack) Top() int {
	if len(this.data) == 0 {
		return -1
	}

	return this.data[len(this.data)-1]
}


func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */