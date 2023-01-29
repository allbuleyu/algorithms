package prob0091

import "testing"

func Test_numDecodings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"x", args{s: "12"}, 2},
		{"x", args{s: "127"}, 2},
		{"x", args{s: "126"}, 3},
		{"x", args{s: "06"}, 0},
		{"x", args{s: "200016"}, 0},
		{"x", args{s: "1123"}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDecodings(tt.args.s); got != tt.want {
				t.Errorf("numDecodings() = %v, want %v", got, tt.want)
			}
		})
	}
}
