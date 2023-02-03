package prob0256

import "testing"

func Test_minCost(t *testing.T) {
	type args struct {
		costs [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"x", args{
			[][]int{{17, 2, 17}, {16, 16, 5}, {14, 3, 19}},
		}, 10},
		{"x", args{
			[][]int{{7, 6, 2}},
		}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCost(tt.args.costs); got != tt.want {
				t.Errorf("minCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
