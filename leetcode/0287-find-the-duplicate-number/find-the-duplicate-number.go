package prob0287

func findDuplicate(nums []int) int {
	return twoPointer(nums)
}

// 快慢指针,其实还不是很懂
func twoPointer(nums []int) int {
	var slow, fast int

	// 因为nums 值 1-n之间,0不在循环考虑范围内
	slow = nums[nums[0]]
	fast = nums[nums[nums[0]]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	// 找到循环入口
	var p1, p2 int
	p1 = slow
	p2 = nums[0]
	for p1 != p2 {
		p1 = nums[p1]
		p2 = nums[p2]
	}

	return p1
}