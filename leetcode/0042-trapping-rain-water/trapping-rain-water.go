package prob0042

func trap(height []int) int {
	return helpStackOptimize(height)
}

// 还有dp和twopointers 的解法

// mono stack
func helpStackOptimize(h []int) int {
	stack := make([]int, 0)
	ans := 0

	for i := 0; i < len(h); i++ {
		// 这里不需要考虑 h[i] == stack.top的情况,因为minH-curH == 0 结果不会有改变
		for len(stack) > 0 && h[stack[len(stack)-1]] < h[i] {

			// 当前高度
			curH := h[stack[len(stack)-1]]
			stack = stack[:len(stack)-1:cap(stack)]

			if len(stack) == 0 {
				break
			}

			// 左边墙高度的索引
			pIndex := stack[len(stack)-1]

			// 两边的墙取高度低的算雨水面积
			minH := min(h[pIndex], h[i])

			//
			ans += (minH-curH) * (i-pIndex-1)
		}


		stack = append(stack, i)
	}

	return ans
}

func helpStack(h []int) int {
	stack := make([]int, 0)
	ans := 0

	for i := 0; i < len(h); i++ {
		if len(stack) > 0 && h[stack[len(stack)-1]] <= h[i] {
			if h[stack[len(stack)-1]] == h[i] {
				stack[len(stack)-1] = i
				continue
			}

			for h[stack[len(stack)-1]] <= h[i] {

				// 当前高度
				curH := h[stack[len(stack)-1]]
				stack = stack[:len(stack)-1:cap(stack)]

				if len(stack) == 0 {
					break
				}

				// 左边墙高度的索引
				pIndex := stack[len(stack)-1]

				// 两边的墙取高度低的算雨水面积
				minH := min(h[pIndex], h[i])

				//
				ans += (minH-curH) * (i-pIndex-1)
			}
		}

		stack = append(stack, i)
	}

	return ans
}

func min(x, y int) int {
	if x > y {
		return y
	}

	return x
}