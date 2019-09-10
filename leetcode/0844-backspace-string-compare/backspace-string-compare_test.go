package prob0844

import (
	"testing"

	"fmt"
	"github.com/stretchr/testify/assert"
)

func Test_normal(t *testing.T)  {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		s, t string
		ans bool
	}{
		//[10], [5], [2], [6], [10, 5], [5, 2], [2, 6], [5, 2, 6]
		{"", "", true},
		{"ab#c", "ad#c", true},
		{"a##c", "#a#c", true},
		{"a#c", "b", false},

	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, twoPointers(tc.s, tc.t), "输入:%v", tc)
	}
}