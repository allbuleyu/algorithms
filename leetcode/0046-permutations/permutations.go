package prob0046


func permute(nums []int) [][]int {
	return recursion(nums)
}

func recursion(nums []int) [][]int {
	res := make([][]int, 0)
	cur := make([]int, 0)

	helpRecursion(&res, nums, &cur)

	return res
}

// 1,2,3 通过for 和把当前的切片缩短,可以替换1,2,3的出现位置
func helpRecursion(res *[][]int, nums []int, cur *[]int) {
	if len(*cur) == len(nums) {
		t := make([]int, len(nums))
		copy(t, *cur)

		// t := *cur	导致结果都为最后一次执行的结果,找一下原因
		*res = append(*res, t)
	}else {
		for i := 0; i < len(nums); i++ {
			if contains(*cur, nums[i]) {
				continue
			}
			*cur = append(*cur, nums[i])
			helpRecursion(res, nums, cur)
			*cur = (*cur)[:len(*cur)-1]
		}
	}
}

func contains(nums []int, num int) bool {
	for _, t := range nums {
		if num == t {
			return true
		}
	}

	return false
}


