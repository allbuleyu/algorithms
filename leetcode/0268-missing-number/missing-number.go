package prob0268

func missingNumber(nums []int) int {
	return solutionBitMap(nums)
}

func solution1(nums []int) int {
	res := len(nums)
	for i := range nums {
		res -= nums[i] - i
	}
	
	return res
}

func solutionBitMap(nums []int) int {
	res := len(nums)
	for i := range nums {
		res ^= i ^ nums[i]
	}

	return res
}

// 数学公式
// 1->100 的数的和 可以算 n*(n+1)/2
func solutionFormula(nums []int) int {
	n := len(nums)
	expectedSum := n*(n+1)/2
	actualSum := 0
	for i := range nums {
		actualSum += nums[i]
	}

	return expectedSum-actualSum
}