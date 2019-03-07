package prob0415

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_prob0415(t *testing.T)  {
	ast := assert.New(t)

	tcs := []struct {
		input string
		input1 string
		ans string
	}{
		{"1", "9", "10"},
		{"123", "123", "246"},
		{"9", "9", "18"},
		{"0", "0", "0"},
		{"9", "99", "108"},

	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, addStrings(tc.input, tc.input1), "输入:%v", tc)
	}
}