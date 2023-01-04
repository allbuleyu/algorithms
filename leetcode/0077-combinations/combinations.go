package prob0077

var combinations [][]int
var cache []int

func combine(n int, k int) [][]int {
	combinations = make([][]int, 0)
	cache = make([]int, 0)
	backtrack(n, k, 1)
	return combinations
}

func backtrack(n, k, begin int) {
	if len(cache) == k {
		cur := make([]int, k)
		copy(cur, cache)
		combinations = append(combinations, cur)
		return
	}

	for i := begin; i <= n; i++ {
		cache = append(cache, i)
		backtrack(n, k, i+1)
		cache = cache[:len(cache)-1]
	}
}
