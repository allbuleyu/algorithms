package sort

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_mergeSortTopDown(t *testing.T) {
	ast := assert.New(t)

	// test case
	tcs := []struct {
		input []int
		ans   []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 3, 4, 2, 1}, []int{1, 2, 3, 4, 5}},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		mergeSortTopDown(tc.input)
		ast.Equal(tc.ans, tc.input, "输入:%v", tc)
	}
}
