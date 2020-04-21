package challenge

import (
	"fmt"
	"github.com/allbuleyu/algorithms/kit"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_iteration(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		input []int
		ans []int
	}{
		// TODO: Add test cases.
		{
			[]int{8,5,1,7,10,12},
			[]int{8,5,1,7,10,12},
		},
	}
	for _, tc := range tests {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, kit.PreOrderBST(iteration(tc.input)), "输入:%v", tc)
	}
}