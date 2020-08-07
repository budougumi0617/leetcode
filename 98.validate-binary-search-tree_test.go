package main

import "testing"

func Test_isValidBST(t *testing.T) {
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
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			want: true,
		},
		{
			name: "example2",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val:   1,
						Left:  &TreeNode{Val: 3},
						Right: &TreeNode{Val: 6},
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
			},
			want: false,
		},
		{
			name: "wrong_answer1",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 1,
					},
				},
			},
			want: false,
		},
		{
			name: "wrong_answer2",
			args: args{
				root: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 1,
					},
				},
			},
			want: false,
		},
		{
			name: "wrong_answer3",
			args: args{
				root: &TreeNode{
					Val: 10,
					Left: &TreeNode{
						Val: 5,
					},
					Right: &TreeNode{
						Val:   15,
						Left:  &TreeNode{Val: 6},
						Right: &TreeNode{Val: 20},
					},
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
