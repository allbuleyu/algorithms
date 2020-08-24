package prob0556

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_nextGreaterElement(t *testing.T) {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input int
		ans int
	}{
		{230241,230412},
		{12,21},
		{21,-1},
		{9876,-1},
		{9876123,9876132},
		{12443322, 13222344},
		{1999999999, -1},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, nextGreaterElement(tc.input), "输入:%v", tc)
	}
}