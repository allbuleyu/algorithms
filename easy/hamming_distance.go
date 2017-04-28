// The Hamming distance between two integers is the number of positions at which the corresponding bits are different.

// Given two integers x and y, calculate the Hamming distance.

// Note:
// 0 ≤ x, y < 231.

// Example:

// Input: x = 1, y = 4

// Output: 2

// Explanation:
// 1   (0 0 0 1)
// 4   (0 1 0 0)
//        ↑   ↑

// The above arrows point to positions where the corresponding bits are different.

package main

import (
	"fmt"
	// "strconv"
)

func main() {
	x := 1
	y := 4

	fmt.Println((x ^ y), hammingDistance(x, y), hammingDistance1(x, y))

	fmt.Printf("%b %b \n", (x ^ y), (x^y)>>1)

	fmt.Println((y << 1) & (x ^ y))
}

func hammingDistance(x int, y int) int {
	r := x ^ y
	c := 0
	for ; r != 0; c++ {
		r &= (r - 1)
	}

	return c
}

// http://www.cnblogs.com/graphics/archive/2010/06/21/1752421.html
func hammingDistance1(x int, y int) int {
	r := x ^ y
	tmp := r - ((r >> 1) & 033333333333) - ((r >> 2) & 011111111111)
	return ((tmp + (tmp >> 3)) & 030707070707) % 63
}
