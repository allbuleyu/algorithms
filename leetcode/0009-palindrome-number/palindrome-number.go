package prob0009


func isPalindrome(x int) bool {
	if x < 0 || (x % 10 == 0 && x != 0) {
		return false
	}

	var tmpx, y, tmpy int
	tmpx = x
	for tmpx > 0 {
		tmpy = tmpx % 10
		tmpx = tmpx / 10
		y = y * 10 + tmpy
	}

	if x != y {
		return false
	}

	return true
}

// 最快速的解决方案 (运行所有test case 所用时间最少)
func isPalindrome1(x int) bool {
	if (x < 0 || (x % 10 == 0 && x != 0)){
		return false
	}

	// 这一步精简,nice
	revertedNumber := 0
	for(x > revertedNumber){
		revertedNumber = revertedNumber * 10 + x % 10
		x = x / 10
	}

	return x == revertedNumber || x == revertedNumber / 10
}