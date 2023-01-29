package prob1335

import "testing"

func Test_minDifficulty(t *testing.T) {
	type args struct {
		jobDifficulty []int
		d             int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"x", args{
			jobDifficulty: []int{3, 2, 1},
			d:             2,
		}, 4},
		{"xy", args{
			jobDifficulty: []int{3, 2, 1},
			d:             1,
		}, 3},
		{"x1", args{
			jobDifficulty: []int{3, 2, 1, 3, 2, 1},
			d:             2,
		}, 4},
		{"x2", args{
			jobDifficulty: []int{9},
			d:             2,
		}, -1},
		{"x3", args{
			jobDifficulty: []int{1, 1, 1},
			d:             3,
		}, 3},
		{"w", args{
			jobDifficulty: []int{11, 111, 22, 222, 33, 333, 44, 444},
			d:             6,
		}, 843},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDifficulty(tt.args.jobDifficulty, tt.args.d); got != tt.want {
				t.Errorf("minDifficulty() = %v, want %v", got, tt.want)
			}
		})
	}
}
