package kit

import (
	"testing"
)

func Test_findKthLargest(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"xxx", args{[]int{3, 2, 1, 5, 6, 4}, 2}, 5},
		{"xxx1", args{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4}, 4},
		{"xxx2", args{[]int{-1, 2, 0}, 1}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargest(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}
