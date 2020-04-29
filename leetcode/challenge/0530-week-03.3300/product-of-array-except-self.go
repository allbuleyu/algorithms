package _530_week_03_3300


//Given an array nums of n integers where n > 1,  return an array output such that output[i] is equal to the product of all the elements of nums except nums[i].
//
//Example:
//
//Input:  [1,2,3,4]
//Output: [24,12,8,6]
//Constraint: It's guaranteed that the product of the elements of any prefix or suffix of the array (including the whole array) fits in a 32 bit integer.
//
//Note: Please solve it without division and in O(n).
//
//Follow up:
//Could you solve it with constant space complexity? (The output array does not count as extra space for the purpose of space complexity analysis.)

func productExceptSelf(nums []int) []int {
	return help2(nums)
}

func help1(nums []int) []int {
	l := make([]int, len(nums))
	r := make([]int, len(nums))

	l[0] = 1
	r[len(nums) - 1] = 1
	for i := 1; i < len(nums); i++ {
		l[i] = l[i-1] * nums[i-1]
	}

	for i := len(nums) - 2; i >= 0; i-- {
		r[i] = r[i+1] * nums[i+1]
	}

	for i := 0; i < len(nums); i++ {
		l[i] *= r[i]
	}

	return l
}

// help1 的优化版本
func help2(nums []int) []int {
	l := 1
	r := 1

	ans := make([]int, len(nums))

	ans[0] = 1
	for i := 1; i < len(nums); i++ {
		l *= nums[i-1]
		ans[i] = l
	}

	ans[len(nums)-1] = l
	for i := len(nums)-2; i >= 0; i-- {
		r *= nums[i+1]
		ans[i] *= r
	}

	return ans
}
