package main

import (
	"reflect"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	// 3 -> 2 -> 0 -> 4 -> 2 -> 0 ....
	exOneList := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}
	exOneWant := exOneList.Next
	if exOneWant.Val != 2 {
		t.Fatalf("exOneWant.Val is not 2, got %d", exOneWant.Val)
	}
	exOneList.Next.Next.Next.Next = exOneWant

	// 1 -> 2 -> 1 -> 2 ...
	exTwoList := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}
	exTwoWant := exTwoList
	if exTwoWant.Val != 1 {
		t.Fatalf("exTwoWant.Val is not 1, got %d", exTwoWant.Val)
	}
	exOneList.Next.Next = exTwoWant

	tests := [...]struct {
		name string
		head *ListNode
		want *ListNode
	}{
		{
			name: "example1",
			head: exOneList,
			want: exOneWant,
		},
		{
			name: "example2",
			head: exTwoList,
			want: exTwoWant,
		},
		{
			name: "example3",
			head: &ListNode{Val: 1},
			want: nil,
		},
		{
			name: "nil",
			head: nil,
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
