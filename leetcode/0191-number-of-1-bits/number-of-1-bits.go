package prob0191

func hammingWeight(num uint32) int {
	return solution3(num)
}

func solution1(num uint32) int {
	n := 0
	for num != 0 {
		tmp := num << 1
		if tmp < num {
			n++
		}

		num = tmp
	}

	return n
}

func solution2(num uint32) int {
	var mask uint32 = 1
	bits := 0

	for i := 0; i < 32; i++ {
		if (num & mask) != 0 {
			bits++
		}
		mask = mask << 1
	}

	return bits
}

func solution3(num uint32) int {
	bits := 0

	for num != 0 {
		bits++
		num = num & (num-1)
	}

	return bits
}