package sort

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_heapSort(t *testing.T) {
	ast := assert.New(t)

	// test case
	type args struct {
		input []int
	}
	tcs := []struct {
		name string
		args args
		ans  []int
	}{
		{"xxxx", args{[]int{9, 8, 2, 1, 4, 5, 6, 7, 3}}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		HeapSort(tc.args.input)
		ast.Equal(tc.ans, tc.args.input, "输入:%v", tc)
	}
}
