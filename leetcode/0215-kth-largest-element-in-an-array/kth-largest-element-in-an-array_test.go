package prob0215

import "testing"

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
		{"x", args{
			nums: []int{5, 6, 7, 9, 12, 13, 8},
			k:    4,
		}, 8},
		{"x", args{
			nums: []int{5, 6, 7, 9, 12, 13, 8},
			k:    1,
		}, 13},
		{"x", args{
			nums: []int{1},
			k:    1,
		}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargest(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}
