package prob0047


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
			[]int{1,1,2},
			[][]int{
				[]int{1,1,2},
				[]int{1,2,1},
				[]int{2,1,1},
			},
		},

	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, permuteUnique(tc.input), "输入:%v", tc)
	}
}