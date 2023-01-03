package prob0203

import (
	"algorithms/kit"
	"testing"

	"fmt"
	"github.com/stretchr/testify/assert"
)

func Test_prb0014(t *testing.T) {
	ast := assert.New(t)

	// test case
	tcs := []struct {
		input []int
		v     int
		ans   []int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 6, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5, 6}, 1, []int{2, 3, 4, 5, 6}},
		{[]int{1, 2, 6}, 6, []int{1, 2}},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		actual := kit.TransferNodes(removeElements(kit.CreateNodes(tc.input), tc.v))
		ast.Equal(tc.ans, actual, "输入:%v", tc)
	}
}
