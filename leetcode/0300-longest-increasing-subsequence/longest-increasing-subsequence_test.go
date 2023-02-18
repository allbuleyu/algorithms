package prob0300

import "testing"

func Test_lengthOfLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"x", args{[]int{10, 9, 2, 5, 3, 7, 101, 18}}, 4},
		{"x1", args{[]int{0, 1, 2, 3}}, 4},
		{"x2", args{[]int{0, 0, 0, 0}}, 1},
		{"x3", args{[]int{7, 6, 5, 4}}, 1},
		{"x3", args{[]int{2, 5, 1, 4}}, 2},
		{"x3", args{[]int{2, 5, 3, 4}}, 3},
		{"x3", args{[]int{0}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
