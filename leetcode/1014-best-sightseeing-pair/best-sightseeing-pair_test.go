package prob1014

import (
	"testing"

	"fmt"
	"github.com/stretchr/testify/assert"
)

func Test_maxScoreSightseeingPair(t *testing.T)  {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input []int

		ans int
	}{
		{[]int{8,1,5,2,6}, 11},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, maxScoreSightseeingPair(tc.input), "输入:%v", tc)
	}
}