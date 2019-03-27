package prob0090

import (
	"testing"

	"fmt"
	"github.com/stretchr/testify/assert"
)

func Test_prb0078(t *testing.T)  {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input []int
		ans [][]int
	}{
		{
			[]int{1,2,2},
			[][]int{
				[]int{},
				[]int{2},
				[]int{1},
				[]int{1,2,2},
				[]int{1,2},
				[]int{2,2},
			},
		},

	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, subsetsWithDup(tc.input), "输入:%v", tc)
	}
}