package sort

func ShellSort(nums []int) {
	n := len(nums)
	h := 1

	// 1, 3, 13, 40, 121, 364, 1093
	for h < n/3 {
		h = 3*h + 1
	}

	for h >= 1 {
		for i := h; i < n; i++ {
			for j := i; j >= h && less(nums, j, j-h); j -= h {
				exchange(nums, j, j-h)
			}
		}
		h /= 3
	}
}
