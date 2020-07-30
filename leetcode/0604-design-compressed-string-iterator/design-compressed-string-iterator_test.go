package prob0604

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_1(t *testing.T) {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input string
		ans string
	}{
		{"L1e2t1C1o1d1e1", "LeetCode"},
		{"L1e10t1C1o1d11e1", "LeeeeeeeeeetCoddddddddddde"},
		{"L1", "L"},
		{"", ""},

	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		c := Constructor(tc.input)
		b := make([]byte, len(tc.ans))
		for i := 0; i < len(tc.ans); i++ {
			b[i] = c.Next()
		}

		ast.Equal(tc.ans, string(b), "输入:%v", tc)
	}
}