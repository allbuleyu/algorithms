package prob0066

func plusOne(digits []int) []int {
	n := len(digits)

	for i := n-1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] += 1

			return digits
		}else {
			digits[i] = 0
		}
	}

	target := make([]int, 1)
	target[0] = 1

	target = append(target, digits...)
	return target
}