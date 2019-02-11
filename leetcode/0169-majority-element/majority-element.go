package prob0169

func majorityElement(nums []int) int {

	return divideAndConquer(nums)
}

func divideAndConquer(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	var left int
	left = len(nums)/2

	var eleLeft, eleRight int
	eleLeft = divideAndConquer(nums[:left])
	eleRight = divideAndConquer(nums[left+1:])

	if eleLeft != eleRight {
		var count int
		for _, v := range nums {
			if v == eleLeft {
				count++
			}
		}

		if count > len(nums)/2 {
			return eleLeft
		}

		return eleRight
	}

	return eleLeft
}

func mapSolution(nums []int) int {
	m := make(map[int]int)
	for i := range nums {
		//if _, ok := m[nums[i]];!ok {
		//	m[nums[i]] = 1
		//}
		m[nums[i]]++
	}

	var max, res int
	for e, n := range m {
		if n > max {
			max = n
			res = e
		}
	}

	return res
}