package prob0084

import (
	"math/rand"
	"testing"
	"time"
)

func Test_largestRectangleAreaOld(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{"1", []int{2, 1, 5, 6, 2, 3}, 10},
		{"1", []int{2, 1, 5, 6, 2, 1}, 10},
		{"1", []int{2, 1, 5, 6, 2, 5, 6}, 10},
		{"2", []int{2, 4}, 4},
		{"3", []int{1, 2, 3, 4, 5}, 9},
		{"4", []int{5, 1}, 5},
		{"5", []int{1, 5, 1}, 5},
		{"6", []int{1, 2, 1}, 3},
		{"7", []int{2, 2, 1, 1, 1}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := withStack(tt.args); got != tt.want {
				t.Errorf("largestRectangleAreaOld() = %v, want %v", got, tt.want)
			}
		})
	}
}

func generateTestData(n int) []int {
	rand.Seed(time.Now().UnixNano())
	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = rand.Intn(n)
	}

	return data
}

func Benchmark_withStack(b *testing.B) {
	data := generateTestData(1000)
	for i := 0; i < b.N; i++ {
		withStack(data)
	}
}

func Benchmark_largestRectangleAreaOld(b *testing.B) {
	data := generateTestData(1000)
	for i := 0; i < b.N; i++ {
		largestRectangleAreaOld(data)
	}
}
