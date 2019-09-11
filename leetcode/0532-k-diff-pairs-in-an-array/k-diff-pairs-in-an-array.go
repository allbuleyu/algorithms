package prob0532

import "sort"

func findPairs(nums []int, k int) int {
	return twoPointers(nums, k)
}

// two pointers
//int ans = 0;
//    Arrays.sort(nums);
//    for (int i = 0, j = 0; i < nums.length; i++) {
//        for (j = Math.max(j, i + 1); j < nums.length && (long) nums[j] - nums[i] < k; j++) ;
//        if (j < nums.length && (long) nums[j] - nums[i] == k) ans++;
//        while (i + 1 < nums.length && nums[i] == nums[i + 1]) i++;
//    }
//    return ans;
func twoPointers(nums []int, k int) int {
	sort.Ints(nums)
	ans := 0
	n := len(nums)
	for i, j := 0,0; i < n; i++ {
		// 没懂这步啥意思
		for j = max(j, i+1); j < n && nums[j] - nums[i] < k; j++ {

		}

		if j < n && nums[j] - nums[i] == k {
			ans++
		}

		for i+1 < n && nums[i] == nums[i+1] {
			i++
		}
	}

	return ans
}

// hash table so easy
func solution2(nums []int, k int) int {
	ans := 0

	//m := map[int]int{}
	for i := 0; i < len(nums); i++ {

	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
