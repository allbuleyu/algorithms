package prob0503

// 单调栈
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	stack := make([]int, 0)
	ans := make([]int, n)


	for i := 2 * n -1; i >= 0; i-- {
		ri := i % n
		for len(stack) > 0 && nums[stack[len(stack)-1]] <= nums[ri] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			ans[ri] = -1
		}else {
			ans[ri] = nums[stack[len(stack)-1]]
		}

		stack = append(stack, ri)
	}

	return ans
}