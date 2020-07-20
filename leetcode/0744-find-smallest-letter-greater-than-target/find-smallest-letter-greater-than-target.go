package prob0744

func nextGreatestLetter(letters []byte, target byte) byte {
	return binSrch(letters, target)
}

func binSrch(letters []byte, target byte) byte {
	left, right := 0, len(letters) -1
	for left < right {
		mid := left + (right-left)/2
		if letters[mid] <= target {
			left = mid+1
		}else {
			right = mid
		}
	}

	//     if letters[right] > target {
	//         return letters[right]
	//     }

	//     return letters[0]
	return letters[left%len(letters)]
}