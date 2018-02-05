package easy


func getSum(a int, b int) int {
	if b == 0 {
		return a
	}

	sum := a ^ b          // 两个数的和(没有进位)
	carry := (a & b) << 1 // 进位

	return getSum(sum, carry)
}
