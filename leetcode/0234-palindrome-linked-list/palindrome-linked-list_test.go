package prob0234


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

		ans bool
	}{
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 2, 1}, true},
		{[]int{1,1}, true},
		{[]int{1}, true},
		{[]int{1, 2, 2, 1}, true},
		{[]int{1, 2}, false},

	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		actual := isPalindrome(kit.CreateNodes(tc.input))
		ast.Equal(tc.ans, actual, "输入:%v", tc)
	}
}