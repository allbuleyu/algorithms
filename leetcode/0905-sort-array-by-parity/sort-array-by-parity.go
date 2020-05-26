package prob0905


// Given an array A of non-negative integers, return an array consisting of all the even elements of A, followed by all the odd elements of A.
//
//You may return any answer array that satisfies this condition.
//
//
//
//Example 1:
//
//Input: [3,1,2,4]
//Output: [2,4,3,1]
//The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.
//
//
//Note:
//
//1 <= A.length <= 5000
//0 <= A[i] <= 5000

func sortArrayByParity(A []int) []int {
	return twoPointers(A)
}

// move zeros 那道题一样的解法.
func twoPointers(nums []int) []int {
	lastEvenIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] % 2 == 0 {
			nums[i], nums[lastEvenIndex] = nums[lastEvenIndex], nums[i]
			lastEvenIndex++
		}
	}

	return nums
}