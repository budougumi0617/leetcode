package main

import "testing"

func Test_isSymmetric(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example1",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 3},
						Right: &TreeNode{Val: 4},
					},
					Right: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 4},
						Right: &TreeNode{Val: 3},
					},
				},
			},
			want: true,
		},
		{
			name: "example2",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   2,
						Right: &TreeNode{Val: 3},
					},
					Right: &TreeNode{
						Val:   2,
						Right: &TreeNode{Val: 3},
					},
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric(tt.args.root); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
