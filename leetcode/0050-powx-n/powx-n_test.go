package prob0050

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_myPow(t *testing.T) {

	ast := assert.New(t)

	// test case
	tcs := []struct{
		input float64
		n int
		ans float64
	}{
		{2, 2, 4},
		{2, -2, 0.25},
		{2, 3, 8},
		{2, 4, 16},
		{2, 5, 32},
	}


	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, myPow(tc.input, tc.n), "输入:%v", tc)
	}
}