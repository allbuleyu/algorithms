package prob0084

// https://leetcode.com/problems/largest-rectangle-in-histogram/description/
// Given n non-negative integers representing the histogram's bar height where the width of each bar is 1, find the area of largest rectangle in the histogram.
//
// Above is a histogram where width of each bar is 1, given height = [2,1,5,6,2,3].
//
// The largest rectangle is shown in the shaded area, which has area = 10 unit.
//
// For example,
// Given heights = [2,1,5,6,2,3],
// return 10.
// next challenge 85
func largestRectangleAreaOld(heights []int) int {
	n := len(heights)
	lessFromLeft := make([]int, n)
	lessFromRight := make([]int, n)
	lessFromLeft[0] = -1
	lessFromRight[n-1] = n

	for i := 1; i < n; i++ {
		p := i - 1
		for p >= 0 && heights[p] >= heights[i] {
			p = lessFromLeft[p]
		}

		lessFromLeft[i] = p
	}

	for i := n - 2; i >= 0; i-- {
		p := i + 1
		for p < n && heights[p] >= heights[i] {
			p = lessFromRight[p]
		}

		lessFromRight[i] = p
	}

	maxArea := 0
	for i := 0; i < n; i++ {
		maxArea = max(maxArea, (lessFromRight[i]-lessFromLeft[i]-1)*heights[i])
	}

	return maxArea
}

// stack 代表单调栈,最后一个元素出栈了,代表了到达当前的floor,所以要乘以 当前遍历到的位置i.
func withStack(heights []int) int {
	stack := make([]int, 0)
	maxArea := 0
	n := len(heights)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if len(stack) == 0 {
				maxArea = max(maxArea, heights[j]*i)
			} else {
				maxArea = max(maxArea, heights[j]*(i-stack[len(stack)-1]-1))
			}
		}

		stack = append(stack, i)
	}

	for len(stack) > 0 {
		j := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if len(stack) == 0 {
			maxArea = max(maxArea, heights[j]*n)
		} else {
			maxArea = max(maxArea, heights[j]*(n-stack[len(stack)-1]-1))
		}
	}

	return maxArea
}

func bruteForce(heights []int) int {
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		minHeight := heights[i]
		maxArea = max(maxArea, heights[i])
		for j := i + 1; j < len(heights); j++ {
			minHeight = min(minHeight, heights[j])
			maxArea = max(maxArea, calculateArea(i, j, minHeight, minHeight))
		}
	}

	return maxArea
}

func calculateArea(x0, x1, y0, y1 int) int {
	return (x1 - x0 + 1) * min(y1, y0)
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func largestRectangleArea(heights []int) int {
	lessLen := 1
	maxArea := 0
	slice := make([]int, 0)
	for i := range heights {
		if i == 0 {
			slice = append(slice, heights[i])
			continue
		}
	Outer:
		for j := len(slice) - 1; j >= 0; j-- {

			if slice[j] > heights[i] {
				max := slice[j] * lessLen
				lessLen++
				slice = slice[:j]

				if max > maxArea {
					maxArea = max
				}

			} else if slice[j] <= heights[i] {

				if lessLen == 1 {
					slice = append(slice, heights[i])
					break Outer
				}

			}
		}

		if lessLen != 1 {
			for k := 0; k < lessLen; k++ {
				slice = append(slice, heights[i])

			}
			lessLen = 1
		}
	}

	for i := 0; i < len(slice); i++ {
		max := slice[i] * (len(slice) - i)
		if max > maxArea {
			maxArea = max
		}
	}

	return maxArea
}

func getSlice(slice *[]int, nums []int) {
	if len(*slice) == 0 {
		return
	}
}

func LargestRectangleArea(heights []int) int {
	return largestRectangleArea(heights)
}
