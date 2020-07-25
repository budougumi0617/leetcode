package main

import (
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "true",
			args: args{s: "A man, a plan, a canal: Panama"},
			want: true,
		},
		{
			name: "false",
			args: args{s: "race a car"},
			want: false,
		},
		{
			name: "false2",
			args: args{s: "0P"},
			want: false,
		},
		{
			name: "no_ward_is_true",
			args: args{s: " "},
			want: true,
		},
		{
			name: "one_word_is_true",
			args: args{s: "a."},
			want: true,
		},
		{
			name: "no_alphanumeric",
			args: args{s: ".,"},
			want: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
