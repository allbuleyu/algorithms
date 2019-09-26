package prob0189

func rotate(nums []int, k int)  {
	k = k % len(nums)
	reverseArr(nums, k)
}

func extraArr(nums []int, k int) {
	ans := make([]int, 0)

	ans = append(ans, nums[len(nums)-k:]...)
	ans = append(ans, nums[:len(nums)-k]...)

	for i := 0; i < len(nums); i++ {
		nums[i] = ans[i]
	}
}

func reverseArr(nums []int, k int) {
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j;  {
		nums[i], nums[j] = nums[j], nums[i]

		i++
		j--
	}
}