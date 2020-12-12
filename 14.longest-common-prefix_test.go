package main

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	tests := [...]struct {
		name string
		strs []string
		want string
	}{
		{name: "Example1", strs: []string{"flower", "flow", "flight"}, want: "fl"},
		{name: "Example2", strs: []string{"dog", "racecar", "car"}, want: ""},
		{name: "Wrong1", strs: []string{"ab", "a"}, want: "a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %q, want %q", got, tt.want)
			}
		})
	}
}
