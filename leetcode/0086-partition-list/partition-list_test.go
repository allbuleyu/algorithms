package prob0086


import (
	"algorithms/kit"
	"testing"

	"github.com/stretchr/testify/assert"
	"fmt"
)

func Test_prb00206(t *testing.T)  {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input []int
		x int
		ans []int
	}{
		{[]int{1, 4,3,2,5,2}, 3, []int{1,2,2,4,3,5}},
		{[]int{1,2,3}, 0, []int{1,2,3}},
		{[]int{1,2,3}, 4, []int{1,2,3}},

	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		actual := kit.TransferNodes(partition(kit.CreateNodes(tc.input),tc.x))
		ast.Equal(tc.ans, actual, "输入:%v", tc)
	}
}