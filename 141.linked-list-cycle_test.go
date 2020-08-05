package main

import "testing"

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
		pos  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example",
			args: args{
				head: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 0,
							Next: &ListNode{
								Val: -4,
							},
						},
					},
				},
				pos: 1,
			},
			want: true,
		},
		{
			name: "example3",
			args: args{
				head: &ListNode{
					Val: 1,
				},
				pos: -1,
			},
		},
		{
			name: "empty",
			args: args{
				pos: -1,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.pos != -1 {
				l := findLast(tt.args.head)
				p := findPos(tt.args.head, tt.args.pos)
				l.Next = p
			}
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func findLast(head *ListNode) *ListNode {
	last := head
	for last.Next != nil {
		last = last.Next
	}
	return last
}

func findPos(head *ListNode, pos int) *ListNode {
	if pos == -1 {
		return nil
	}
	p := head
	for i := 0; i < pos; i++ {
		p = p.Next
	}
	return p
}
