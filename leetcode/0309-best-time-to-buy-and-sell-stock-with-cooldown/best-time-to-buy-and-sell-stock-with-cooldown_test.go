package prob0309

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"x", args{prices: []int{1, 2, 3, 0, 2}}, 3},
		{"x", args{prices: []int{1}}, 0},
		{"x", args{prices: []int{1, 0}}, 0},
		{"x", args{prices: []int{1, 3, 2, 8, 4, 9}}, 8},
		{"x", args{prices: []int{1, 3, 2, 9, 4, 9}}, 8},
		{"x", args{prices: []int{1, 5, 2, 8, 4, 9}}, 9},
		{"x", args{prices: []int{3, 2, 1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
