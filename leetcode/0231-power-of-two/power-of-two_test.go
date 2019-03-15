package prob0231

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)
var tcs = []struct{
	input int
	ans bool
}{
	{8, true},
	{4, true},
	{2, true},
	{1, true},
	{0, false},
	{5, false},

}

func Test_missingNumber(t *testing.T) {
	ast := assert.New(t)

	// test case
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, isPowerOfTwo(tc.input), "输入:%v", tc)
	}
}

func Benchmark_solution1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			solution1(tc.input)
		}
	}
}

func Benchmark_solution2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			solution2(tc.input)
		}
	}
}