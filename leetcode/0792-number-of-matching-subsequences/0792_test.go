package prob0792

import (
	"testing"
)

func Test_numMatchingSubseq(t *testing.T) {
	type args struct {
		s     string
		words []string
	}
	tests := []struct {
		name    string
		args    args
		wantNum int
	}{
		{"x", args{
			s:     "agbcde",
			words: []string{"a", "bb", "acd", "ace"},
		}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNum := numMatchingSubseq(tt.args.s, tt.args.words); gotNum != tt.wantNum {
				t.Errorf("numMatchingSubseq() = %v, want %v", gotNum, tt.wantNum)
			}
		})
	}
}

func Test_Map(t *testing.T) {
	m := map[int]int{}

	for i := 0; i < 10; i++ {
		m[i] = i
	}

	t.Log(m[111])
}
