package prob0190

func reverseBits(num uint32) uint32 {
	return solutionOptimize(num)
}

// 注释部分为其它方法
func solution1(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		//if num & 1 == 1 {
		//	res = (res << 1) + 1
		//}else {
		//	res = res << 1
		//}

		res = (res << 1) | (num & 1)
		num >>= 1
	}

	return res
}

func solution2(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		// 各种位运算
		res = (res << 1) | ((num >> uint32(i)) & 1)
		//res = (num >> uint32(i) & 1) << (31 - uint32(i)) | res
	}

	return res
}

// https://leetcode.com/problems/reverse-bits/discuss/54741/O(1)-bit-operation-C%2B%2B-solution-(8ms)
// 论坛得票最高的解答
func solutionOptimize(n uint32) uint32 {
	n = (n >> 16) | (n << 16)
	//n = ((n & 0xffffffff) >> 16) | ((n & 0xffffffff) << 16)
	n = ((n & 0xff00ff00) >> 8) | ((n & 0x00ff00ff) << 8)
	n = ((n & 0xf0f0f0f0) >> 4) | ((n & 0x0f0f0f0f) << 4)
	n = ((n & 0xcccccccc) >> 2) | ((n & 0x33333333) << 2)
	n = ((n & 0xaaaaaaaa) >> 1) | ((n & 0x55555555) << 1)

	return n
}