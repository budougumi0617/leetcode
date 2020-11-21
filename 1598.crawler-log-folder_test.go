package main

import "testing"

func Test_minOperations(t *testing.T) {
	tests := []struct {
		name string
		logs []string
		want int
	}{
		{name: "Example1", logs: []string{"d1/", "d2/", "../", "d21/", "./"}, want: 2},
		{name: "Example2", logs: []string{"d1/", "d2/", "./", "d3/", "../", "d31/"}, want: 3},
		{name: "Example3", logs: []string{"d1/", "../", "../", "../"}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.logs); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
