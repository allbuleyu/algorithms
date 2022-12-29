package sort

import (
	"math/rand"
	"time"
)

func QuickSort(nums []int) {
	RandomQuickSort(nums, 0, len(nums)-1)
}

func RandomQuickSort(nums []int, lo, hi int) {
	rand.Seed(time.Now().UnixNano())
	randomQuickSort(nums, lo, hi)
}

func randomQuickSort(nums []int, lo, hi int) {
	if hi <= lo {
		return
	}

	i := rand.Int()%(hi-lo+1) + lo
	swap(nums, i, hi)
	i = partition(nums, lo, hi)
	randomQuickSort(nums, lo, i-1)
	randomQuickSort(nums, i+1, hi)
}

func quickSort(nums []int, lo, hi int) {
	// 这里的条件只能是hi <= lo, 通过长度为1 的数组来模拟,可得出结论.
	if hi <= lo {
		return
	}

	i := partition(nums, lo, hi)
	quickSort(nums, lo, i-1)
	quickSort(nums, i+1, hi)
}

func partition(nums []int, lo, hi int) int {
	x := nums[hi]
	i := lo - 1

	for j := lo; j < hi; j++ {
		if nums[j] <= x {
			i++
			swap(nums, i, j)
		}
	}

	// 为什么不用这个,因为这样i,k的距离就会非常远,
	// 很可能会用不到高速缓存.
	//k := j - 1
	//for i < k {
	//	if nums[i] <= x {
	//		i++
	//	} else {
	//		swap(nums, i, k)
	//		k--
	//	}
	//}

	swap(nums, i+1, hi)
	return i + 1
}

// func partition(nums []int, lo, hi int) int {
// 	x := nums[lo]
// 	i, j := lo, hi

// 	for {
// 		for nums[i+1] < x {
// 			i++
// 			if i == hi {
// 				break
// 			}
// 		}

// 		for x < nums[j] {
// 			j--
// 		}

// 		if j <= i {
// 			break
// 		}

// 		swap(nums, i, j)
// 	}

// 	swap(nums, j, lo)
// 	return j
// }

// func less(nums []int, i, j int) bool {
// 	return nums[i] <= nums[j]
// }
