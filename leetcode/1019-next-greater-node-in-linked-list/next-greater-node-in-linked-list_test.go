package prob1019

import (
	"fmt"
	"algorithms/kit"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_oddEvenList(t *testing.T)  {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input []int
		ans []int
	}{
		{[]int{2,1,5}, []int{5,5,0}},
		{[]int{2,7,4,3,5}, []int{7,0,5,5,0}},
		{[]int{1,7,5,1,9,2,5,1}, []int{7,9,9,9,0,5,0,0}},
		{[]int{1}, []int{0}},

	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, nextLargerNodes(kit.CreateNodes(tc.input)), "输入:%v", tc)
	}
}