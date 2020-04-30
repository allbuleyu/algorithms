package _530_week_03_3306

//(This problem is an interactive problem.)
//
//A binary matrix means that all elements are 0 or 1. For each individual row of the matrix, this row is sorted in non-decreasing order.
//
//Given a row-sorted binary matrix binaryMatrix, return leftmost column index(0-indexed) with at least a 1 in it. If such index doesn't exist, return -1.
//
//You can't access the Binary Matrix directly.  You may only access the matrix using a BinaryMatrix interface:
//
//BinaryMatrix.get(row, col) returns the element of the matrix at index (row, col) (0-indexed).
//BinaryMatrix.dimensions() returns a list of 2 elements [rows, cols], which means the matrix is rows * cols.
//Submissions making more than 1000 calls to BinaryMatrix.get will be judged Wrong Answer.  Also, any solutions that attempt to circumvent the judge will result in disqualification.
//
//For custom testing purposes you're given the binary matrix mat as input in the following four examples. You will not have access the binary matrix directly.
//
//
//
//
//
//
//
//Example 1:
//
//
//
//Input: mat = [[0,0],[1,1]]
//Output: 0
//Example 2:
//
//
//
//Input: mat = [[0,0],[0,1]]
//Output: 1
//Example 3:
//
//
//
//Input: mat = [[0,0],[0,0]]
//Output: -1
//Example 4:
//
//
//
//Input: mat = [[0,0,0,1],[0,0,1,1],[0,1,1,1]]
//Output: 1
//
//
//Constraints:
//
//rows == mat.length
//cols == mat[i].length
//1 <= rows, cols <= 100
//mat[i][j] is either 0 or 1.
//mat[i] is sorted in a non-decreasing way.

type BinaryMatrix interface {
	Get(int, int) int
	Dimensions() []int
}
/**
 * // This is the BinaryMatrix's API interface.
 * // You should not implement it, or speculate about its implementation
 * type BinaryMatrix struct {
 *     Get(int, int) int
 *     Dimensions() []int
 * }
 */

func leftMostColumnWithOne(binaryMatrix BinaryMatrix) int {
	dim := binaryMatrix.Dimensions()
	if len(dim) < 2 {
		return -1
	}
	m, n := dim[0], dim[1]
	minIndex := -1
	for i := 0; i < m; i++ {
		tmp := get(binaryMatrix, i, n)
		if tmp != -1 {
			if minIndex == -1 {
				minIndex = tmp
			}else {
				minIndex = min(minIndex, tmp)
			}
		}
	}

	return minIndex
}

func get(binaryMatrix BinaryMatrix, m, n int) int {
	l := 0
	r := n - 1

	index := -1
	for l <= r {
		mid := (l+r)/2
		if binaryMatrix.Get(m, mid) == 0 {
			l = mid + 1
		}else {
			r = mid - 1
			index = mid
		}
	}

	return index
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func findFirstIndex(nums []int) int {
	l := 0
	r := len(nums) - 1

	index := -1
	for l <= r {
		mid := (l+r)/2
		if nums[mid] == 0 {
			l = mid + 1
		}

		if nums[mid] == 1 {
			r = mid - 1
			index = mid
		}

	}

	return index
}
