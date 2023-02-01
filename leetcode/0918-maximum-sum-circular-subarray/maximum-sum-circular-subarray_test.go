package prob0918

import "testing"

func Test_maxSubarraySumCircular(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"x", args{nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}}, 6},
		{"x", args{nums: []int{5, 4, -1, 7, 8}}, 23},
		{"x", args{nums: []int{1}}, 1},
		{"x", args{nums: []int{-1, -2, -3}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubarraySumCircular(tt.args.nums); got != tt.want {
				t.Errorf("maxSubarraySumCircular() = %v, want %v", got, tt.want)
			}
		})
	}
}
