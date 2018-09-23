package prob0263

func isUgly(num int) bool {
	if num <= 0 {
		return false
	}

	for num > 6 {
		if num % 2 != 0 && num % 3 != 0 && num % 5 != 0 {
			return false
		}

		if num % 2 == 0 {
			num = num / 2
		}

		if num %3 == 0 {
			num = num / 3
		}

		if num %5 == 0 {
			num = num / 5
		}
	}

	return true
}