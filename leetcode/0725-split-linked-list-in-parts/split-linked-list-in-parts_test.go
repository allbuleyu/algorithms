package prob0725


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
		input2 int
		ans [][]int
	}{
		{[]int{1, 2, 3}, 5, [][]int{
				[]int{1},[]int{2},[]int{3},[]int{},[]int{},
			},
		},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3, [][]int{
			[]int{1, 2, 3, 4},[]int{5, 6, 7},[]int{8, 9, 10},
			},
		},
		{[]int{}, 3, [][]int{
			[]int{},[]int{},[]int{},
		},
		},

	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		actual := make([][]int, tc.input2)
		for i, a := range splitListToParts(kit.CreateNodes(tc.input), tc.input2) {
			actual[i] = kit.TransferNodes(a)
		}

		ast.Equal(tc.ans, actual, "输入:%v", tc)
	}
}