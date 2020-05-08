package week1_3320

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_firstUniqChar(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		input string
		ans int
	}{
		// TODO: Add test cases.
		{"loveleetcode", 2},
	}

	for _, tc := range tests {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, firstUniqChar(tc.input), "输入:%v", tc)
	}
}