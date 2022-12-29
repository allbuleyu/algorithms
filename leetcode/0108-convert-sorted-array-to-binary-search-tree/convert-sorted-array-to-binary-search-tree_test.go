package prob0108

import (
	"fmt"
	"algorithms/kit"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sortedArrayToBST(t *testing.T) {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input []int
		preOrder []int
	}{
		{[]int{-10,-3,0,5,9}, []int{0, -10, -3, 5, 9}},


	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		root := sortedArrayToBST(tc.input)
		ast.Equal(tc.input, kit.InOrder(root), "输入:%v", tc)
		ast.Equal(tc.preOrder, kit.PreOrder(root), "输入:%v", tc)
	}
}