package prob0119

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getRow(t *testing.T)  {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input int
		ans []int
	}{
		//{[]int{10, 5, 2, 6},100,8},
		{4, []int{1,4,6,4,1}},
		{3, []int{1,3,3,1}},

	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, getRow(tc.input), "输入:%v", tc)
	}
}