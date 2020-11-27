package main

import "testing"

func Test_rob(t *testing.T) {
	tests := [...]struct {
		name string
		nums []int
		want int
	}{
		{name: "Example1", nums: []int{1, 2, 3, 1}, want: 4},
		{name: "Example2", nums: []int{2, 7, 9, 3, 1}, want: 12},
		{name: "Complex1", nums: []int{100, 2, 4, 2000, 5}, want: 2100},
		{name: "Complex2", nums: []int{10, 2, 150, 200, 150}, want: 310},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
