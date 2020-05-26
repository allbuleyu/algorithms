package prob0283


// Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.
//
//Example:
//
//Input: [0,1,0,3,12]
//Output: [1,3,12,0,0]
//Note:
//
//You must do this in-place without making a copy of the array.
//Minimize the total number of operations.

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