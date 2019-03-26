package prob0078

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
			[]int{1,2,3},
			[][]int{
				[]int{},
				[]int{1},
				[]int{2},
				[]int{3},
				[]int{1,2},
				[]int{1,3},
				[]int{2,3},
				[]int{1,2,3},
			},
		},

	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, subsets(tc.input), "输入:%v", tc)
	}
}