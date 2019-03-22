package prob0476

func findComplement(num int) int {
	var x int
	x = ^0

	for (x & num) > 0 {
		x <<= 1
	}

	x = ^x

	return x^num

}