package prob0092

import (
	"algorithms/kit"
	"testing"

	"github.com/stretchr/testify/assert"
	"fmt"
)

func Test_prb0234(t *testing.T)  {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input []int
		m,n int
		ans []int
	}{
		{[]int{1, 2}, 1,2, []int{2,1}},
		{[]int{1, 2, 3, 4, 5}, 2,4, []int{1, 4, 3,2,5}},

	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		actual := kit.TransferNodes(reverseBetween(kit.CreateNodes(tc.input), tc.m, tc.n))
		ast.Equal(tc.ans, actual, "输入:%v", tc)
	}
}