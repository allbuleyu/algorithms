package prob0154

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findMin(t *testing.T) {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input []int
		ans int
	}{
		{[]int{2,2,2,0,1},0},
		{[]int{2,2,2,2},2},
		{[]int{2,1,1,1,1,1},1},
		{[]int{1,3,5},1},
	}


	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, findMin(tc.input), "输入:%v", tc)
	}

}