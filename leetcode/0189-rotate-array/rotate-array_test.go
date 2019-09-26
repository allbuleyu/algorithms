package prob0189


import (
	"testing"

	"fmt"
	"github.com/stretchr/testify/assert"
)

func Test_rotate(t *testing.T)  {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input []int
		k int
		ans []int
	}{
		{[]int{1,2,3,4,5,6,7}, 3, []int{5,6,7,1,2,3,4}},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		rotate(tc.input, tc.k)
		ast.Equal(tc.ans, tc.input, "输入:%v", tc)
	}
}