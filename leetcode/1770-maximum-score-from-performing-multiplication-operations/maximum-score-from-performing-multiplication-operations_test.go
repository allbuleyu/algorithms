package prob1770

import "testing"

func Test_maximumScore(t *testing.T) {
	type args struct {
		nums        []int
		multipliers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{
			nums:        []int{1, 2, 3},
			multipliers: []int{1, 2, 3},
		}, 14},
		{"2", args{
			nums:        []int{-5, -3, -3, -2, 7, 1},
			multipliers: []int{-10, -5, 3, 4, 6},
		}, 102},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumScore(tt.args.nums, tt.args.multipliers); got != tt.want {
				t.Errorf("maximumScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Array(t *testing.T) {
	arr := newArrAndReturun(t)
	arrPointer := newArrAndReturunPointer(t)

	t.Logf("external func arr pointer:%p", &arr)
	t.Logf("external func arrPointer pointer:%p", arrPointer)
}

func newArrAndReturun(t *testing.T) []int {
	n := 10
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = i
	}

	t.Logf("internal func arr pointer:%p", &res)
	return res
}

func newArrAndReturunPointer(t *testing.T) *[]int {
	n := 10
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = i
	}

	t.Logf("internal func arrPointer pointer:%p", &res)
	return &res
}
