package prob0740

import "testing"

func Test_deleteAndEarn(t *testing.T) {

	tests := []struct {
		name string
		args []int
		want int
	}{
		{"1", []int{1, 2}, 2},
		{"1x", []int{3, 4, 2}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteAndEarn(tt.args); got != tt.want {
				t.Errorf("deleteAndEarn() = %v, want %v", got, tt.want)
			}
		})
	}
}
