package prob0012

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_intToRoman(t *testing.T) {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input string
		ans int
	}{
		{"III", 3},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXCIV", 1994},

	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.input, intToRoman(tc.ans), "输入:%v", tc)
	}
}