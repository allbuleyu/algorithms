package week2_3327


//You are given a sorted array consisting of only integers where every element appears exactly twice, except for one element which appears exactly once. Find this single element that appears only once.
//
//
//
//Example 1:
//
//Input: [1,1,2,3,3,4,4,8,8]
//Output: 2
//Example 2:
//
//Input: [3,3,7,7,10,11,11]
//Output: 10
//
//
//Note: Your solution should run in O(log n) time and O(1) space.

func singleNonDuplicate(nums []int) int {
	return helpBinarySearch(nums)
}

func help1(nums []int) int {
	res := nums[len(nums)-1]
	repeatNum := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			repeatNum++
			continue
		}

		if repeatNum == 0 {
			res = nums[i-1]
			break
		}

		repeatNum = 0
	}

	return res
}

func helpBinarySearch(nums []int) int {
	var l, r int
	l = 0
	r = len(nums)-1

	for l < r {
		mid := l + (r-l)/2
		if mid % 2 == 1 {
			mid--
		}

		if nums[mid] != nums[mid+1] {
			r = mid
		}else {
			l = mid+2
		}
	}

	return nums[l]
}
