package prob0011

func maxArea(height []int) int {
	var left,right, maxArea int
	left = 0
	right = len(height)-1


	for left < right {
		maxArea = max(maxArea, (right-left)*min(height[left], height[right]))

		if height[left] < height[right] {
			left++
		}else {
			right--
		}

	}

	return maxArea
}

func max(x, y int) int {
	if x < y {
		return y
	}

	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}