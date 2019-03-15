package prob0231

func isPowerOfTwo(n int) bool {
	return solution2(n)
}

func solution1(n int) bool {
	if n <= 0 {
		return false
	}

	for n > 1 {
		if n & 1 == 1 {
			return false
		}

		n >>= 1
	}

	return true
}

func solution2(n int) bool {
	//if n <= 0 {
	//	return false
	//}
	//
	//if n & (n-1) != 0 {
	//	return false
	//}
	//
	//return true

	// 更精简
	//return n > 0 && (n&(-n)) == n
	return n > 0 && (n & (n-1)) == 0
}