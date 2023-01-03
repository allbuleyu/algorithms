package sort

var aux []int

func MergeSortDownUp(nums []int) {
	n := len(nums)
	aux = make([]int, n)
	for sz := 1; sz < n; sz += sz {
		for lo := 0; lo < n-sz; lo += sz + sz {
			merge(nums, lo, lo+sz-1, min(lo+sz+sz-1, n-1))
		}
	}
}

func MergeSortUpDown(nums []int) {
	aux = make([]int, len(nums))
	mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, lo, hi int) {
	if lo >= hi {
		return
	}
	mid := lo + (hi-lo)/2
	mergeSort(nums, lo, mid)
	mergeSort(nums, mid+1, hi)
	merge(nums, lo, mid, hi)
}

func merge(nums []int, lo, mid, hi int) {
	i, j := lo, mid+1
	for k := lo; k <= hi; k++ {
		aux[k] = nums[k]
	}

	for k := lo; k <= hi; k++ {
		if i > mid {
			nums[k] = aux[j]
			j++
		} else if j > hi {
			nums[k] = aux[i]
			i++
		} else if less(aux, i, j) {
			nums[k] = aux[i]
			i++
		} else {
			nums[k] = aux[j]
			j++
		}
	}
}

func mergeSortTopDown(nums []int) {
	helpMergeSortTopDown(nums, 0, len(nums)-1)
}

func helpMergeSortTopDown(nums []int, start, end int) {
	if start == end {
		return
	}

	mid := (end-start)/2 + start

	helpMergeSortTopDown(nums, start, mid)
	helpMergeSortTopDown(nums, mid+1, end)

	combine(nums, start, mid, end)
}

func combine(nums []int, start, mid, end int) {
	for mid < end {
		if nums[start] <= nums[mid] {
			start++
		} else {
			swap(nums, start, mid)
			start++
			mid++
		}
	}
}
