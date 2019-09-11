package prob0560

func subarraySum(nums []int, k int) int {

}

func bruteForce(nums []int, k int) int {
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