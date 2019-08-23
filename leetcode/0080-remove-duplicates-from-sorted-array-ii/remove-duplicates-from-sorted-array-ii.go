package prob0080

func removeDuplicates(nums []int) int {
	return discussSolution(nums)
}


//int i = 0;
//    for (int n : nums)
//        if (i < 2 || n > nums[i-2])
//            nums[i++] = n;
//    return i;
func discussSolution(nums []int) int {
	i := 0

	for _, v := range nums {
		if i < 2 || v > nums[i-2] {
			nums[i] = v
			i++
		}
	}

	return i
}

func towPointers(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var i, j, ctn int
	ctn = 1
	for j = 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]

			ctn = 1
		}else {
			ctn++
			if ctn <= 2 {
				i++
				nums[i] = nums[j]
			}
		}
	}

	return i+1
}