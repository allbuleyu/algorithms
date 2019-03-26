package prob0040

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int, 0)
	cur := make([]int, 0)

	helper(&res, &cur, candidates, target, 0)
	return  res
}

func helper(res *[][]int, cur *[]int, candidates []int, target, start int)  {
	if target < 0 {
		return
	}

	if target == 0 {
		c := make([]int, len(*cur))
		copy(c, *cur)
		*res = append(*res, c)
	}else {
		for i := start; i < len(candidates) && candidates[i] <= target; i++ {
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}

			*cur = append(*cur, candidates[i])
			helper(res, cur, candidates, target-candidates[i], i+1)
			*cur = (*cur)[:len(*cur)-1]
		}

	}
}