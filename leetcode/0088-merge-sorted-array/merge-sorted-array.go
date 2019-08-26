package prob0088

func merge(nums1 []int, m int, nums2 []int, n int)  {
	len1 := len(nums1)-1

	//for m > 0 || n > 0 {
	//	if m == 0 {
	//		for n > 0 {
	//			nums1[len1] = nums2[n-1]
	//			n--
	//			len1--
	//		}
	//		break
	//	}
	//
	//	if n == 0 {
	//		break
	//	}
	//
	//	if nums1[m-1] > nums2[n-1] {
	//		nums1[len1] = nums1[m-1]
	//		len1--
	//		m--
	//	}else {
	//		nums1[len1] = nums2[n-1]
	//		n--
	//		len1--
	//	}
	//}


	// 优化
	n--
	m--
	for m >= 0 && n >= 0 {
		if nums1[m] > nums2[n] {
			nums1[len1] = nums1[m]
			len1--
			m--
		}else {
			nums1[len1] = nums2[n]
			n--
			len1--
		}
	}

	for n >= 0 {
		nums1[len1] = nums2[n]
		n--
		len1--
	}
}