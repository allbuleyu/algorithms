package prob1299

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_replaceElements(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		input []int
		ans []int
	}{
		// TODO: Add test cases.

		{[]int{17,18,5,4,6,1}, []int{18,6,6,6,1,-1}},
	}

	for _, tc := range tests {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, replaceElements(tc.input), "输入:%v", tc)
	}
}