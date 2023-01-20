package prob0139

import "testing"

func Test_wordBreak(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"x", args{
			s:        "leetcode",
			wordDict: []string{"leet", "code"},
		}, true},
		{"x", args{
			s:        "leetcodee",
			wordDict: []string{"leet", "code"},
		}, false},
		{"x", args{
			s:        "applepenapple",
			wordDict: []string{"apple", "pen"},
		}, true},
		{"x", args{
			s:        "applepenapplepen",
			wordDict: []string{"apple", "pen"},
		}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordBreak(tt.args.s, tt.args.wordDict); got != tt.want {
				t.Errorf("wordBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}
