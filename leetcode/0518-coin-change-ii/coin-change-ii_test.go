package prob0518

import "testing"

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test", args{
			coins:  []int{1, 2},
			amount: 3,
		}, 2},
		{"x", args{
			coins:  []int{1, 2, 5},
			amount: 5,
		}, 4},
		{"x", args{
			coins:  []int{2},
			amount: 3,
		}, 0},
		{"x", args{
			coins:  []int{10},
			amount: 10,
		}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := change(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
