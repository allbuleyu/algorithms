package prob0739

func dailyTemperatures(T []int) []int {
	// mono stack optimize
	stack := make([]int, 0)
	ans := make([]int, len(T))
	for i := len(T)-1; i >= 0; i-- {
		for len(stack) > 0 && T[stack[len(stack)-1]] <= T[i] {
			stack = stack[:len(stack)-1:cap(stack)]
		}

		if len(stack) == 0 {
			ans[i] = 0
		}else {
			ans[i] = stack[len(stack)-1]-i
		}

		stack = append(stack, i)
	}

	return ans
}

func helpMonoStack(t []int) []int {
	stack := make([]int, 0)
	hs := make(map[int]int)

	for i := 0; i < len(t); i++ {
		for len(stack) > 0 && t[stack[len(stack)-1]] < t[i] {
			m := len(stack)
			hs[stack[m-1]] = i
			stack = stack[:m-1:cap(stack)]
		}

		stack = append(stack, i)
	}

	ans := make([]int, len(t))
	for i := 0; i < len(t); i++ {
		k, ok := hs[i]
		if !ok {
			ans[i] = 0
		}else {
			ans[i] = k-i
		}
	}

	return ans
}
