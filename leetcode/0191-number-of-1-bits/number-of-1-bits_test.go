package prob0191

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	input uint32
	ans int
}{
	{1, 1},
	{2, 1},
	{11, 3},
	{128, 1},
	{8, 1},
	{3, 2},
}

func TestHammingWeight(t *testing.T) {
	var i uint32 = 8
	fmt.Println(i<<1)

	ast := assert.New(t)



	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, hammingWeight(tc.input),"输入:%v", tc)
	}
}

func Benchmark_hammingWeight(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			hammingWeight(tc.input)
		}
	}
}