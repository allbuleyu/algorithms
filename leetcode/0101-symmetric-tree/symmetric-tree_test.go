package prob0101

import (
	"fmt"
	"github.com/allbuleyu/algorithms/kit"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isSymmetric(t *testing.T) {

	tcs := []struct{
		input []int
		ans bool
	}{
		{
			[]int{1,2,2,3,4,4,3},
			true,
		},
		{
			[]int{1,2,2,kit.Null,3,kit.Null,3},
			false,
		},
	}


	ast := assert.New(t)
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, isSymmetric(kit.NewTrees(tc.input).Root), "输入:%v", tc)
	}
}