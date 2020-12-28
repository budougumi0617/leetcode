package main

import (
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	tests := []struct {
		name       string
		head, want *ListNode
	}{
		{
			name: "Example1",
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
					},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
				},
			},
		},
		{
			name: "Example2",
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 3,
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
		},
		{
			name: "nil",
			head: nil,
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := deleteDuplicates(tt.head)
			want := tt.want
			for {
				if want.Next == nil && got.Next == nil {
					break
				}
				if got.Val != want.Val {
					t.Fatalf("want val %d, but got %d", want.Val, got.Val)
				}
				got, want = got.Next, want.Next
			}
		})
	}
}
