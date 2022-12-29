package prob0206

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

		ans []int
	}{
		{[]int{1, 2, 3}, []int{3,2,1}},
		{[]int{1, 2, 3,4}, []int{4,3,2,1}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		actual := kit.TransferNodes(reverseList(kit.CreateNodes(tc.input)))
		ast.Equal(tc.ans, actual, "输入:%v", tc)
	}
}