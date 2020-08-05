package prob0496

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_nextGreaterElement(t *testing.T) {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input, input2 []int
		ans []int
	}{
		{[]int{4,1,2}, []int{1,3,4,2}, []int{-1,3,-1}},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, nextGreaterElement(tc.input, tc.input2), "输入:%v", tc)
	}
}