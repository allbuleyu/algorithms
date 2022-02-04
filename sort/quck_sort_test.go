package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"xxx", args{[]int{1, 2}}},
		{"xxx", args{[]int{2, 1}}},
		{"xxx", args{[]int{1, 2, 3, 4, 5, 6, 7}}},
		{"xxx", args{[]int{5, 4, 7, 6, 9, 8}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSort(tt.args.nums)

		})
		fmt.Println(tt.args.nums)
	}
}
