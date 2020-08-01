package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_deleteNode(t *testing.T) {
	type args struct {
		nodeVal int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "example",
			args: args{
				nodeVal: 5,
			},
			want: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 9,
					},
				},
			},
		},
		{
			name: "example2",
			args: args{
				nodeVal: 1,
			},
			want: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 9,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 9,
						},
					},
				},
			}
			node := got
			for node.Val != tt.args.nodeVal {
				node = node.Next
			}

			deleteNode(node)
			if d := cmp.Diff(got, tt.want); len(d) != 0 {
				t.Errorf("differs: (-got +want)\n%s", d)
			}
		})
	}
}
