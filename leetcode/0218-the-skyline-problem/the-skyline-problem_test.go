package prob0218

import (
	"reflect"
	"testing"
)

type myTest struct {
	name string
	args [][]int
	want [][]int
}

func generateTests() []myTest {
	tests := []myTest{
		{"x", [][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}}, [][]int{{2, 10}, {3, 15}, {7, 12}, {12, 0}, {15, 10}, {20, 8}, {24, 0}}},
		{"x1", [][]int{{0, 2, 3}, {2, 5, 3}}, [][]int{{0, 3}, {5, 0}}},
	}
	return tests
}

func Test_getSkyline(t *testing.T) {
	tests := generateTests()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSkyline(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSkyline() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bruteForce2(t *testing.T) {
	tests := generateTests()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bruteForce2(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSkyline() = %v, want %v", got, tt.want)
			}
		})
	}
}
