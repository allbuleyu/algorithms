package prob0069

import "math"

func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	return NewtonMethod(x)
}

func binarySearch(x int) int {
	left, right := 2.0, float64(x/2)
	var pivot float64
	for left <= right {
		pivot = left +( right-left)/2

		if pivot*pivot > float64(x) {
			right = pivot-1
		}else if pivot*pivot < float64(x) {
			left = pivot+1
		}

		return int(pivot)
	}

	return int(right)
}

// root fo x = 2 * root of pow(x)/4
// x 的跟 = 2 * (x的平方*1/4) 开根号
func recursionBitShift(x int) int {
	if x < 2 {
		return x
	}

	left := recursionBitShift(x >> 2) << 1
	right := left + 1

	if right * right > x {
		return left
	}

	return right
}

// 最好,应用最广泛的求根的方法
// 牛顿方法.看不懂递推式,先记录下来
func NewtonMethod(x int) int {
	if x < 2 {
		return x
	}

	var x0, x1, fx float64
	x0 = float64(x)
	fx = x0
	x1 = (x0 + fx/x0)/2.0

	for math.Abs(x0-x1) >= 1 {
		x0 = x1
		x1 = (x0 + fx/x0) / 2.0
	}

	return int(x1)
}