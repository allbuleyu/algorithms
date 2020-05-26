package _528_week_01_3286

//https://leetcode.com/explore/featured/card/30-day-leetcoding-challenge/528/week-1/3286/
//Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.
//
//Example:
//
//Input: [0,1,0,3,12]
//Output: [1,3,12,0,0]
//Note:
//
//You must do this in-place without making a copy of the array.
//Minimize the total number of operations.
//   Hide Hint #1
//In-place means we should not be allocating any space for extra array. But we are allowed to modify the existing array. However, as a first step, try coming up with a solution that makes use of additional space. For this problem as well, first apply the idea discussed using an additional array and the in-place solution will pop up eventually.
//   Hide Hint #2
//A two-pointer approach could be helpful here. The idea would be to have one pointer for iterating the array and another pointer that just works on the non-zero elements of the array.
func moveZeroes(nums []int)  {
	twoPointersOptimal(nums)
}

// two pointers
func moveZeroesOptimize(nums []int) {
	var i, j int
	for j = 0; j < len(nums); j++ {
		if nums[j] != 0 {
			nums[i] = nums[j]
			i++
		}
	}

	for i < len(nums) {
		nums[i] = 0
		i++
	}
}

func twoPointers(nums []int) {
	lastNoneZeroIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[lastNoneZeroIndex] = nums[i]
			lastNoneZeroIndex++
		}
	}

	for lastNoneZeroIndex < len(nums) {
		nums[lastNoneZeroIndex]=0
		lastNoneZeroIndex++
	}
}

// 牛逼啊
func twoPointersOptimal(nums []int) {
	lastNoneZeroIndex := 0
	for cur := 0; cur < len(nums); cur++ {
		if nums[cur] != 0 {
			nums[cur], nums[lastNoneZeroIndex] = nums[lastNoneZeroIndex], nums[cur]
			lastNoneZeroIndex++
		}
	}
}
