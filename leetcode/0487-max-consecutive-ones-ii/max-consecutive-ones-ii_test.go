package prob0487

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findMaxConsecutiveOnes(t *testing.T) {
	ast := assert.New(t)

	tcs := []struct{
		input []int
		ans int
	}{
		//{[]int{1,0,1,1,0},4},
		{[]int{1,0,1,1,1},5},
		{[]int{1,0,1,1,0,1,1},5},
		{[]int{1,0,1,1,0,1,1,0},5},
		{[]int{1,1,1,1,1},5},
		{[]int{1},1},
		{[]int{0},1},
		{[]int{},0},

	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, findMaxConsecutiveOnes(tc.input), "输入:%v", tc)
	}
}