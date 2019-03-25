package prob0039

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int, 0)
	cur := make([]int, 0)

	helperOptimize(&res, &cur, candidates, target, 0)
	return  res
}

func helper(res *[][]int, cur, candidates []int, target, start int) {
	if target < 0 {
		return
	}

	if target == 0 {
		*res = append(*res, cur)
	}else {
		for i := start; i < len(candidates) && candidates[i] <= target; i++ {
			cur = append(cur, candidates[i])
			helper(res, cur, candidates, target-candidates[i], i)
			cur = cur[:len(cur)-1:len(cur)-1]	// 把底层数组的容量也改变,要不然会有错误答案
		}
	}
}

// 优化helper方法,使用的内存更少了
func helperOptimize(res *[][]int, cur *[]int, candidates []int, target, start int) {
	if target < 0 {
		return
	}

	if target == 0 {
		c := make([]int, len(*cur))
		copy(c, *cur)
		*res = append(*res, c)
	}else {
		for i := start; i < len(candidates) && candidates[i] <= target; i++ {
			*cur = append(*cur, candidates[i])
			helperOptimize(res, cur, candidates, target-candidates[i], i)
			*cur = (*cur)[:len(*cur)-1]
		}
	}
}