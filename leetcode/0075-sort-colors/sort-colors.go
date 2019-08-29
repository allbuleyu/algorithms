package prob0075

func sortColors(nums []int)  {
	quickSort3(nums)
}

func solution1(nums []int) {
	var n0, n1, n2 int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[n2] = 2
			nums[n1] = 1
			nums[n0] = 0

			n0++
			n1++
			n2++
		}else if nums[i] == 1 {
			nums[n2] = 2
			nums[n1] = 1
			n1++
			n2++
		}else {
			nums[n2] = 2
			n2++
		}
	}
}

// 此方法完全看不懂是啥意思
func solution2(nums []int) {
	var zero, two int
	two = len(nums)-1

	for i := 0; i <= two; i++ {
		if nums[i] == 0 && i > zero {
			nums[i], nums[zero] = nums[zero], nums[i]
			i--
			zero++
		}else if nums[i] == 2 && i < two {

			nums[i], nums[two] = nums[two], nums[i]
			i--
			two--
		}
	}
}

func optimizeSolution2(nums []int) {
	var zero, two int
	two = len(nums)-1

	for i := 0; i <= two; i++ {
		if nums[i] == 0 {
			nums[i], nums[zero] = nums[zero], nums[i]
			zero++
		} else if nums[i] == 2 {
			nums[i], nums[two] = nums[two], nums[i]
			i--
			two--
		}
	}
}

// quick sort three ways
// 快速排序三路排序的思路
func quickSort3(nums []int) {
	var l, r, i int
	r = len(nums)-1
	for i <= r {
		if nums[i] == 0 {
			nums[i], nums[l] = nums[l], nums[i]
			l++
			i++
		}else if nums[i] == 1 {
			i++
		}else {
			nums[r], nums[i] = nums[i], nums[r]
			r--

		}
	}
}