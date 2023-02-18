package prob1171

import (
	"reflect"
	"testing"
)

func Test_removeZeroSumSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"x", args{nums: []int{1, 2, -3, 3, 4}}, []int{1, 2, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeZeroSumSubArray(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeZeroSumSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
