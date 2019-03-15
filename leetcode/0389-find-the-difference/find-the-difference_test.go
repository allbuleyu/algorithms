package prob0389

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_prob0389(t *testing.T) {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input string
		input1 string
		ans byte
	}{
		{"abcd", "abcde", 'e'},
		{"c", "ac", 'a'},
		{"", "a", 'a'},


	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, findTheDifference(tc.input, tc.input1), "输入:%v", tc)
	}
}