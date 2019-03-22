package prob0191

const (
	m1 = 0x55555555
	m2 = 0x33333333
	m4 = 0x0f0f0f0f
	m8 = 0x00ff00ff
	m16 = 0x0000ffff
)

func hammingWeight(num uint32) int {
	return solution4(num)
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

// bench 100000000	        19.5 ns/op
func solution3(num uint32) int {
	bits := 0

	for num != 0 {
		bits++
		num = num & (num-1)
	}

	return bits
}

// bench 100000000	        3.5 ns/op
// https://en.wikipedia.org/wiki/Hamming_weight
func solution4(num uint32) int {
	num = (num & m1) + (num >> 1) & m1		// 计算2位上1的个数
	num = (num & m2) + (num >> 2) & m2		// 计算4位上1的个数
	num = (num & m4) + (num >> 4) & m4		// 计算8位上1的个数
	num = (num & m8) + (num >> 8) & m8		// 计算16位上1的个数
	num = (num & m16) + (num >> 16) & m16		// 计算32位上1的个数

	return int(num)
}