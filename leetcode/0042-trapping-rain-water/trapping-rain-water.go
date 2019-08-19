package prob0042

func trap(height []int) int {
	return bruteForce(height)
}

func twoPointer(height []int) int {
	var ans, left, right, leftMax, rightMax int
	right=len(height)-1

	for left < right {
		if height[left] < height[right] {
			if height[left] > leftMax {
				leftMax = height[left]

			}else {
				ans += leftMax - height[left]
			}

			left++
		}else {
			if height[right] > rightMax {
				rightMax = height[right]
			}else{
				ans += rightMax-height[right]
			}
			right--
		}
	}

	return ans
}

func bruteForce(height []int) int {
	var ans, maxLeft, maxRight int

	for i := 0; i < len(height); i++ {
		maxLeft=0
		maxRight=0

		for j := i; j > 0; j-- {
			maxLeft = max(maxLeft, height[j])
		}

		for j := i; j < len(height); j++ {
			maxRight = max(maxRight, height[j])
		}

		ans += min(maxLeft, maxRight) - height[i]
	}

	return ans
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