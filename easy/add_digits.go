package easy

import (
	"fmt"
	"math"
)


// Given a non-negative integer num, repeatedly add all its digits until the result has only one digit.

// For example:

// Given num = 38, the process is like: 3 + 8 = 11, 1 + 1 = 2. Since 2 has only one digit, return it.

// Follow up:
// Could you do it without any loop/recursion in O(1) runtime?
func addDigits(num int) int {
	if num%10 == num || num <= 0 {
		return num
	}

	sum := 0
	for num != 0 {
		sum += num % 10
		num = num / 10
	}

	return addDigits(sum)
}

func addDigits1(num int) int {

	return ((num - 1) % 9) + 1
}

// Write an algorithm to determine if a number is "happy".

// A happy number is a number defined by the following process: Starting with any positive integer, replace the number by the sum of the squares of its digits, and repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1. Those numbers for which this process ends in 1 are happy numbers.

// Example: 19 is a happy number

// 1*1 + 9*9 = 82
// 8*8 + 2*2 = 68
// 6*6 + 8*8 = 100
// 1*1 + 0*0 + 0*0 = 1
// 2 4 16 37 58 89 125 30 9 81 65 61 37

func isHappy(n int) bool {
	x := n
	m := make(map[int]int)
	m[n] = 1
	for x > 0 {
		sum := addHappy(x)

		if sum == 1 {
			return true
		}
		x = sum
		// fmt.Println(m, sum)
		if _, ok := m[sum]; ok {
			return false
		}
		m[sum] = 1
		fmt.Println(m, sum)
	}

	return true
}

func addHappy(num int) int {
	sum := 0
	for num != 0 {
		sum += (num % 10) * (num % 10)
		num /= 10
	}

	return sum
}

func isUgly(num int) bool {
	if num <= 0 {
		return false
	}

	// for num >= 2 {
	// 	switch {
	// 	case num%2 == 0:
	// 		num /= 2
	// 	case num%3 == 0:
	// 		num /= 3
	// 	case num%5 == 0:
	// 		num /= 5
	// 	default:
	// 		return false
	// 	}
	// }

	arr := [3]int{2, 3, 5}
	for _, i := range arr {

		for num%i == 0 {
			num /= i
		}

		if num == 1 {
			return true
		}
	}

	return false
}

func nthUglyNumber(n int) int {
	if n <= 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	var ugly2, ugly3, ugly5 int
	var result [1690]float64
	result[0] = 1

	for i := 1; i < n; i++ {
		result[i] = math.Min(result[ugly2]*2, math.Min(result[ugly3]*3, result[ugly5]*5))

		if result[i] == result[ugly2]*2 {
			ugly2++
		}

		if result[i] == result[ugly3]*3 {
			ugly3++
		}

		if result[i] == result[ugly5]*5 {
			ugly5++
		}
	}

	return int(result[n-1])
}
