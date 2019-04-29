package prob0047

func permuteUnique(nums []int) [][]int {
	return recursion(nums)
}

func recursion(nums []int) [][]int {
	res := make([][]int, 0)
	cur := make([]int, 0)
	isUsed := make([]bool, len(nums))
	helpRecursion(&res, nums, &cur, &isUsed)

	return res
}

func helpRecursion(res *[][]int, nums []int, cur *[]int, isUsed *[]bool) {

}