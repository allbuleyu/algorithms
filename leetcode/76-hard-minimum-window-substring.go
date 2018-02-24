package leetcode

//https://leetcode.com/problems/minimum-window-substring/description/
//Given a string S and a string T, find the minimum window in S which will contain all the characters in T in complexity O(n).
//
//For example,
//S = "ADOBECODEBANC"
//T = "ABC"
//Minimum window is "BANC".
//
//Note:
//If there is no such window in S that covers all characters in T, return the empty string "".
//
//If there are multiple such windows, you are guaranteed that there will always be only one unique minimum window in S.
//next challenge 30(h),209,239(h),632(h)
// 移动窗口
func minWindow1(s string, t string) string {
	var hs [58]int // 'z'-'A'
	var count int
	for _, v := range t {
		hs[v-'A']++
		count++
	}

	var start int
	var minStr string

	for i := 0; i < len(s); i++ {
		v := s[i]
		if count != 0 {
			hs[v-'A']--
			if hs[v-'A'] >= 0 {
				count--
			}
		}


		if count == 0 {
			if minStr == "" || len(minStr) > len(s[start:i+1]) {
				minStr = s[start:i+1]
			}
		}

		// 移动窗口最重要的一环
		if i-start >= len(t) && count == 0 {
			//hs[s[start]-'A']++
			//if hs[s[start]-'A'] > 0 {
			//	count++
			//	i++ // 此处的意义为,count != 0 所以i需要直接移动到下一位
			//}
			//start++
			//i--
			for count == 0 {
				hs[s[start]-'A']++
				if hs[s[start]-'A'] > 0 {
					count++
				}

				if len(minStr) > len(s[start:i+1]) {
					minStr = s[start:i+1]
				}

				start++
			}

		}
	}

	return minStr
}

// 移动窗口
func minWindow(s string, t string) string {
	var hs [58]int // 'z'-'A'
	var count int
	for _, v := range t {
		hs[v-'A']++
		count++
	}

	var start int
	var minStr string

	for i := 0; i < len(s); i++ {
		v := s[i]

		hs[v-'A']--
		if hs[v-'A'] >= 0 {
			count--
		}

		if count == 0 {
			if minStr == "" || len(minStr) > len(s[start:i+1]) {
				minStr = s[start:i+1]
			}
		}

		// 移动窗口最重要的一环
		for i+1-start >= len(t) && count == 0 {
			hs[s[start]-'A']++
			if hs[s[start]-'A'] > 0 {
				count++
			}

			if len(minStr) > len(s[start:i+1]) {
				minStr = s[start:i+1]
			}

			start++
		}
	}

	return minStr
}

func MinWindow(s string, t string) string {
	return minWindow(s, t)
}