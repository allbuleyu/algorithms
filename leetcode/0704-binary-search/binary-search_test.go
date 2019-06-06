package prob0704

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_search(t *testing.T)  {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input []int
		target int
		ans int
	}{
		{[]int{-1,0,3,5,9,12}, 9, 4,},
		{[]int{-1,0,3,5,9,12}, 2, -1,},
		{[]int{1}, 1, 0,},
		{[]int{1}, 2, -1,},
		{[]int{5}, -5, -1,},

	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, search(tc.input, tc.target), "输入:%v", tc)
	}
}