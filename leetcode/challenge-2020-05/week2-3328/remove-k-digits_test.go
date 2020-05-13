package week2_3328

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_removeKdigits(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		input string
		k int
		ans string
	}{
		// TODO: Add test cases.
		{"1432219", 3, "1219"},
		{"12345", 1, "1234"},
		{"10200", 1, "200"},
		{"10", 2, "0"},
	}

	for _, tc := range tests {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, removeKdigits(tc.input, tc.k), "输入:%v", tc)
	}
}