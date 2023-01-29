package prob0276

import "testing"

func Test_totalWays(t *testing.T) {
	type args struct {
		i int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"x", args{
			i: 3,
			k: 2,
		}, 6},
		{"x", args{
			i: 1,
			k: 1,
		}, 1},
		{"x", args{
			i: 7,
			k: 2,
		}, 42},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalWays(tt.args.i, tt.args.k); got != tt.want {
				t.Errorf("totalWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
