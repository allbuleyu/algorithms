package prob0084


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
// next challenge 85
func largestRectangleAreaOld(heights []int) int {
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
		for j := len(slice) - 1; j >= 0;j-- {

			if slice[j] > heights[i] {
				max := slice[j] * lessLen
				lessLen++
				slice = slice[:j]

				if max > maxArea {
					maxArea = max
				}

			}else if slice[j] <= heights[i]{

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
			lessLen=1
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