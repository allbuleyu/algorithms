package week3_3331

import (
	"fmt"
	"github.com/allbuleyu/algorithms/kit"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_oddEvenList(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		input []int
		ans []int
	}{
		// TODO: Add test cases.
		{[]int{1,2,3,4,5}, []int{1,3,5,2,4}},
		{[]int{2,1,3,5,6,4,7}, []int{2,3,6,7,1,5,4}},
		{[]int{1,2}, []int{1,2}},
		{[]int{1}, []int{1}},
	}

	for _, tc := range tests {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, kit.TransferNodes(oddEvenList(kit.CreateNodes(tc.input))), "输入:%v", tc)
	}
}