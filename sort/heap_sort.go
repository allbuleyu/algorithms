package sort

func HeapSort(nums []int) {
	heapSort(nums)
}

func heapSort(nums []int) {
	n := len(nums)
	for i := n / 2; i >= 0; i-- {
		sink(nums, i, n)
	}

	for i := n - 1; i > 0; i-- {
		exchange(nums, 0, i)
		n--
		sink(nums, 0, n)
	}
}

func sink(nums []int, i, n int) {
	for i*2+1 < n {
		j := i*2 + 1
		if j < n-1 && less(nums, j, j+1) {
			j++
		}

		if less(nums, j, i) {
			break
		}

		exchange(nums, i, j)
		i = j
	}
}
