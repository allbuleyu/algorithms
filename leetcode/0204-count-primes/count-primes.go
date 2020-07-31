package prob0204

func countPrimes(n int) int {
	prime := make([]bool, n)
	for i := 2; i < n; i++ {
		prime[i] = true
	}

	for i := 2; i * i < n; i++ {
		if prime[i] == false {
			continue
		}

		for j := i * i; j < n; j += i {
			prime[j] = false
		}
	}

	ans := 0
	for i := 2; i < n; i++ {
		if prime[i] {
			ans++
		}
	}

	return ans
}


