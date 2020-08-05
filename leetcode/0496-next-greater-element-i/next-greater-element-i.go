package prob0496

// 此处用到单调(递减)栈, 真的是非常优雅的方法.
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	hs := make(map[int]int)
	stack := make([]int, 0)
	ans := make([]int, len(nums1))

	for i := 0; i < len(nums2); i++ {
		for len(stack) > 0 && nums2[i] > stack[len(stack)-1] {
			hs[stack[len(stack)-1]] = nums2[i]
			stack = stack[:len(stack)-1:cap(stack)]
		}

		stack = append(stack, nums2[i])
	}

	for i := 0; i < len(nums1); i++ {
		if v, ok := hs[nums1[i]]; ok {
			ans[i] = v
		}else {
			ans[i] = -1
		}
	}

	return ans
}
