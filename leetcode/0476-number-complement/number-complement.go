package prob0476

func findComplement(num int) int {
	if num == 0 {
		return 1
	}

	res := 0
	var i uint8
	for num != 0 {
		if num & 1 != 1 {
			res += 1 << i
		}
		i++
		num >>= 1
	}

	return res
}