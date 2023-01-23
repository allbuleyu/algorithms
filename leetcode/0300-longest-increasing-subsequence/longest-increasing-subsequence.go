package prob0300

func lengthOfLIS(nums []int) int {
	return buildLIS(nums)
}

// this approach is very graceful
func buildLIS(nums []int) int {
	sub := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] > sub[len(sub)-1] {
			sub = append(sub, nums[i])
		} else {
			//j := 0
			//for nums[i] > sub[j] {
			//	j++
			//}

			// optimize 上面的注释部分.
			j := binarySearch(sub, nums[i])

			sub[j] = nums[i]
		}
	}

	return len(sub)
}

func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

func bottomUp(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	longest := 0
	for i := 0; i < n; i++ {
		longest = max(longest, dp[i])
	}

	return longest
}

var memory []int

func helperUpDown(nums []int) int {
	memory = make([]int, len(nums))
	longest := 0

	upDown(nums, len(nums)-1)

	for i := range memory {
		longest = max(longest, memory[i])
	}
	return longest
}

func upDown(nums []int, start int) int {
	if start == 0 {
		return 1
	}

	if memory[start] > 0 {
		return memory[start]
	}

	longest := 1
	for i := start - 1; i >= 0; i-- {
		if nums[start] > nums[i] {
			longest = max(longest, upDown(nums, i)+1)
		}
	}

	memory[start] = longest
	return memory[start]
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
