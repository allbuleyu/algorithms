package prob1155

import "testing"

func Test_numRollsToTarget(t *testing.T) {
	type args struct {
		n      int
		k      int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"x", args{
			n:      1,
			k:      6,
			target: 3,
		}, 1},
		{"p", args{
			n:      2,
			k:      6,
			target: 7,
		}, 6},
		{"y", args{
			n:      30,
			k:      30,
			target: 500,
		}, 222616187},
		{"y", args{
			n:      10,
			k:      16,
			target: 20,
		}, 92378},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numRollsToTarget(tt.args.n, tt.args.k, tt.args.target); got != tt.want {
				t.Errorf("numRollsToTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
