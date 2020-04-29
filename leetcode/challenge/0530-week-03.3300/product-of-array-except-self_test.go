package _530_week_03_3300

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_productExceptSelf(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		input []int
		ans []int
	}{
		// TODO: Add test cases.
		{[]int{1,2,3,4}, []int{24,12,8,6}},



	}

	for _, tc := range tests {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, productExceptSelf(tc.input), "输入:%v", tc)
	}
}