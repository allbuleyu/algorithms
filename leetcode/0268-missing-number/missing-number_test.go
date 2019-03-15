package prob0268

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_missingNumber(t *testing.T) {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input []int
		ans int
	}{
		{[]int{0,1,3}, 2},
		{[]int{9,6,4,2,3,5,7,0,1}, 8},
		{[]int{0}, 1},

	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, missingNumber(tc.input), "输入:%v", tc)
	}
}

func Benchmark_missingNumber(b *testing.B) {
	tcs := []struct{
		input []int
		ans int
	}{
		{[]int{0,1,3}, 2},
		{[]int{9,6,4,2,3,5,7,0,1}, 8},

	}

	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			missingNumber(tc.input)
		}
	}
}