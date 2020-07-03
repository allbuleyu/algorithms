package prob0031


// 题干的意思是:找到当前序列的下一个序列
// 比如 12345
// 下一个序列就是12354, 刚好比上一个序列大
func nextPermutation(nums []int)  {
	reverseIndex := 0

	i := len(nums) - 1
	j := i
	for ; i > 0; i-- {
		if nums[i] <= nums[i-1] {
			continue
		}

		reverseIndex = i
		for ; j >= i; j-- {
			if nums[j] > nums[i-1] {
				nums[j], nums[i-1] = nums[i-1], nums[j]
				break
			}
		}

		break
	}

	for i,j = reverseIndex, len(nums)-1; i < j; {
		nums[j], nums[i] = nums[i], nums[j]
		i++
		j--
	}
}