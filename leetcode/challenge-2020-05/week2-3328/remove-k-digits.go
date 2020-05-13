package week2_3328

//Given a non-negative integer num represented as a string, remove k digits from the number so that the new number is the smallest possible.
//
//Note:
//The length of num is less than 10002 and will be ≥ k.
//The given num does not contain any leading zero.
//Example 1:
//
//Input: num = "1432219", k = 3
//Output: "1219"
//Explanation: Remove the three digits 4, 3, and 2 to form the new number 1219 which is the smallest.
//Example 2:
//
//Input: num = "10200", k = 1
//Output: "200"
//Explanation: Remove the leading 1 and the number is 200. Note that the output must not contain leading zeroes.
//Example 3:
//
//Input: num = "10", k = 2
//Output: "0"
//Explanation: Remove all the digits from the number and it is left with nothing which is 0.

func removeKdigits(num string, k int) string {
	return helpFunc1(num, k)
}

func helpFunc1(num string, k int) string {
	res := make([]byte, 0)

	n := len(num)
	keep := n - k
	for i := 0; i < n; i++ {
		for k > 0 && len(res) > 0 && res[len(res)-1] > num[i] {
			res = res[:len(res)-1]
			k--
		}

		res = append(res, num[i])
	}

	if len(res) < keep {
		for i := 0; i < keep-len(res); i++ {
			res = append(res, 0)
		}
	}

	if k > 0 {
		res = res[:len(res)-k]
	}

	for len(res) > 0 && res[0] == '0' {
		res = res[1:]
	}

	if len(res) == 0 {
		return "0"
	}

	return string(res)
}