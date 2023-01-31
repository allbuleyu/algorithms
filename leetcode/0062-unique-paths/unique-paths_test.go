package prob0062

import "testing"

func Test_uniquePaths(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"x", args{2, 3}, 3},
		{"x", args{3, 7}, 28},
		{"x", args{1, 1}, 1},
		{"x", args{1, 7}, 1},
		{"x", args{7, 1}, 1},
		{"x", args{7, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
