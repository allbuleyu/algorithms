package _531_week_04_3313


type ListNode struct {
	Val int
	Next *ListNode
}

type FirstUnique struct {
	root *ListNode
}


func Constructor(nums []int) FirstUnique {
	var root *ListNode
	unique := make([]int, 0)
	isUnique := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if _, ok := isUnique[nums[i]]; ok {
			isUnique[nums[i]] = false
			continue
		}

		isUnique[nums[i]] = true
	}

	for i := 0; i < len(nums); i++ {
		if isUnique[nums[i]] == true {
			unique = append(unique, nums[i])
		}
	}

	if len(unique) == 0 {
		root = nil
	}

	root = &ListNode{
		Val:  unique[0],
		Next: nil,
	}
	cur := root
	for i := 1; i < len(unique); i++ {
		next := &ListNode{
			Val: unique[i],
		}

		cur.Next = next
		cur = cur.Next
	}

	return FirstUnique{root: root}
}


func (this *FirstUnique) ShowFirstUnique() int {
	return this.root.Val
}


func (this *FirstUnique) Add(value int)  {
	if this.root == nil {
		this.root = &ListNode{
			Val:  value,
			Next: nil,
		}
		return
	}

	pre := &ListNode{
		Next: this.root,
	}
	cur := pre

	for cur.Next != nil {
		if cur.Next.Val == value {
			cur.Next = cur.Next
		}
	}
}


/**
 * Your FirstUnique object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.ShowFirstUnique();
 * obj.Add(value);
 */
