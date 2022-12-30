package prob0328


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
		ans []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1,3,5,2,4}},
		{[]int{2,1,3,5,6,4,7}, []int{2,3,6,7,1,5,4}},
		{[]int{1, 2, 3, 4}, []int{1,3,2,4}},

	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		actual := kit.TransferNodes(helpStupid(kit.CreateNodes(tc.input)))
		ast.Equal(tc.ans, actual, "输入:%v", tc)
	}
}