package prob0046


import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_permute(t *testing.T) {
	ast := assert.New(t)

	tcs := []struct{
		input []int
		ans [][]int
	}{
		{
			[]int{1,2,3},
			[][]int{
				[]int{1,2,3},
				[]int{1,3,2},
				[]int{2,1,3},
				[]int{2,3,1},
				[]int{3,1,2},
				[]int{3,2,1},
			},
		},

	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, permute(tc.input), "输入:%v", tc)
	}
}