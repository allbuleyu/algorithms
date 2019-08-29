package prob0287

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findDuplicate(t *testing.T) {

	tcs := []struct{
		input []int

		ans int
	}{
		{
			[]int{1,3,4,2,2},2,
		},
		{
			[]int{3,1,3,4,2},3,
		},
	}

	ast := assert.New(t)
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, findDuplicate(tc.input), "输入:%v", tc)
	}
}