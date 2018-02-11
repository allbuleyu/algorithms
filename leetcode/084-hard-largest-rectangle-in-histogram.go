package leetcode

//https://leetcode.com/problems/largest-rectangle-in-histogram/description/
//Given n non-negative integers representing the histogram's bar height where the width of each bar is 1, find the area of largest rectangle in the histogram.
//
//
//Above is a histogram where width of each bar is 1, given height = [2,1,5,6,2,3].
//
//
//The largest rectangle is shown in the shaded area, which has area = 10 unit.
//
//For example,
//Given heights = [2,1,5,6,2,3],
//return 10.

func largestRectangleArea(heights []int) int {
	m := make(map[int]int)
	maxArea := 0
	for i := range heights {
		m[heights[i]] = 1
		for j := i+1; j <= len(heights) -1; j++ {
			if heights[j] >= heights[i] {
				m[heights[i]]++

				if maxArea < heights[i] * m[heights[i]] {
					maxArea = heights[i] * m[heights[i]]
				}
			}else {
				if maxArea < heights[i] * m[heights[i]] {
					maxArea = heights[i] * m[heights[i]]
				}
				m[heights[i]] = 0
				break
			}
		}

		if maxArea < heights[i] * m[heights[i]] {
			maxArea = heights[i] * m[heights[i]]
		}
	}

	return maxArea
}

func LargestRectangleArea(heights []int) int {
	return largestRectangleArea(heights)
}