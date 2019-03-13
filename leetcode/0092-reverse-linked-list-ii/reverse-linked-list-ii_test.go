package prob0092

import (
	"github.com/allbuleyu/algorithms/kit"
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
		{[]int{1, 2, 3, 4}, 2,2, []int{1, 2,3,4}},

	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		actual := reverseBetween(kit.CreateNodes(tc.input), tc.m, tc.n)
		ast.Equal(tc.ans, actual, "输入:%v", tc)
	}
}