package prob0050

func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	return recursion(x, n)
}

func fastPowRecursion(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}

	half := fastPowRecursion(x, n/2)
	half *= half

	// n&1 != 0 等价于 n%2 != 0
	if n%2 != 0 {
		half *= x
	}

	return half
}

func fastPow(x float64, n int) float64 {
	pow := float64(1)
	for n > 0 {
		if n&1 != 0 {
			pow *= x
		}
		x *= x
		n >>= 1
	}

	return pow
}

func recursion(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if x == 0 {
		return 0
	}

	half := recursion(x, n/2)
	half *= half
	if n%2 == 1 {
		half *= x
	}

	return half
}
