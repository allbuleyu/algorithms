package prob0240

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_searchMatrix(t *testing.T) {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input [][]int
		target int
		ans bool
	}{
		{
			[][]int{
				[]int{-1,3},

				//[]int{7,8,9,10,30},

			}, 3, true,
		},
		{
			[][]int{
				[]int{1,2},
				[]int{2,30},

				//[]int{7,8,9,10,30},

			}, 20, false,
		},
		{
			[][]int{
				[]int{1,   4,  7, 11, 15},
				[]int{2,   5,  8, 12, 19},
				[]int{3,   6,  9, 16, 22},
				[]int{10, 13, 14, 17, 24},
				[]int{18, 21, 23, 26, 30},
			}, 30, true,
		},
		{
			[][]int{
				[]int{1,   4,  7, 11, 15},
				[]int{2,   5,  8, 12, 19},
				[]int{3,   6,  9, 16, 22},
				[]int{10, 13, 14, 17, 24},
				[]int{18, 21, 23, 26, 30},
			}, 22, true,
		},

	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, searchMatrix(tc.input, tc.target), "输入:%v", tc)
	}
}