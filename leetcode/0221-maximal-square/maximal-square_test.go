package prob0221

import "testing"

func Test_maximalSquare(t *testing.T) {
	type args struct {
		matrix [][]byte
	}
	tests := []struct {
		name string
		args [][]byte
		want int
	}{
		{"x", [][]byte{{'1', '1', '1'}, {'1', '1', '1'}, {'1', '1', '1'}}, 9},
		{"x1", [][]byte{{'1', '1', '1'}, {'1', '1', '1'}, {'1', '1', '0'}}, 4},
		{"x2", [][]byte{{'0', '1'}, {'1', '0'}}, 1},
		{"x3", [][]byte{{'1'}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalSquare(tt.args); got != tt.want {
				t.Errorf("maximalSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
