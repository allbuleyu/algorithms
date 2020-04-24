package _531_week_04_3307

//Given an array of integers and an integer k, you need to find the total number of continuous subarrays whose sum equals to k.
//
//Example 1:
//Input:nums = [1,1,1], k = 2
//Output: 2
//Note:
//The length of the array is in range [1, 20,000].
//The range of numbers in the array is [-1000, 1000] and the range of the integer k is [-1e7, 1e7].
func subarraySum(nums []int, k int) int {
	return subarraySumHs(nums, k)
}

func subarraySumHs(nums []int, k int) int {
	// Sliding window -- No, contains negative number
	// hashmap + preSum
	/*
	   1. Hashmap<sum[0,i - 1], frequency>
	   2. sum[i, j] = sum[0, j] - sum[0, i - 1]    --> sum[0, i - 1] = sum[0, j] - sum[i, j]
	          k           sum      hashmap-key     -->  hashmap-key  =  sum - k
	   3. now, we have k and sum.
	         As long as we can find a sum[0, i - 1], we then get a valid subarray
	        which is as long as we have the hashmap-key,  we then get a valid subarray
	   4. Why don't map.put(sum[0, i - 1], 1) every time ?
	         if all numbers are positive, this is fine
	         if there exists negative number, there could be preSum frequency > 1
	*/
	hs := make(map[int]int)
	var ans, sum int

	hs[0] = 1
	for _, num := range nums {
		sum += num
		if _, ok := hs[sum-k]; ok {
			ans += hs[sum-k]
		}

		if preSum, ok := hs[sum];ok {
			hs[sum] = preSum+1
		}else {
			hs[sum] = 1
		}
	}

	return ans
}

func solution1(nums []int, k int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]

			if sum == k {
				ans++
			}
		}
	}

	return ans
}