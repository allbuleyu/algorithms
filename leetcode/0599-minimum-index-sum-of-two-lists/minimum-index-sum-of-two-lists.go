package prob0599

func findRestaurant(list1 []string, list2 []string) []string {
	hs := make(map[string]int)
	for i := 0; i < len(list1); i++ {
		hs[list1[i]] = i
	}

	ans := make([]string, 0)
	var minSum, sum int
	minSum = 2000
	for i := 0;  i< len(list2);	i++ {
		if i1, ok := hs[list2[i]];ok {
			sum = i1+i
			if sum < minSum {
				ans = ans[:0]
				ans = append(ans, list2[i])
				minSum = sum
			}else if sum == minSum {
				ans = append(ans, list2[i])
			}
		}
	}

	return ans
}