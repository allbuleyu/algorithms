package prob0022

import (
	"reflect"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	tests := []struct {
		name string
		args int
		want []string
	}{
		// TODO: Add test cases.
		{"1", 1, []string{"()"}},
		{"2", 2, []string{"(())", "()()"}},
		{"3", 3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateParenthesis(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}
