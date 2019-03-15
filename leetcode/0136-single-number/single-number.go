package prob0136


func singleNumber(nums []int) int {
	var res int

	for i := 0; i < len(nums); i++ {
		res ^= nums[i]
	}

	return res
}

// 其它的解法
// hash map, ListNode 就不说了
// 2*(a+b+c)-(a+a+b+b+c) = c	用到了集合
func solution1()  {

}