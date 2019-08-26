package prob0088

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_merge(t *testing.T) {

	tcs := []struct{
		input []int
		input1 []int
		m,n int
		ans []int
	}{
		{
			[]int{1,2,3,0,0,0},[]int{2,5,6},3,3,[]int{1,2,2,3,5,6},
		},
		{
			[]int{6,7,8,0,0,0},[]int{2,5,6},3,3,[]int{2,5,6,6,7,8},
		},
	}

	ast := assert.New(t)
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		merge(tc.input, tc.m, tc.input1,tc.n)
		ast.Equal(tc.ans, tc.input, "输入:%v", tc)
	}
}