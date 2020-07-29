package main

import (
	"testing"
)

func TestMinStack_Pop(t *testing.T) {
	ms := &MinStack{}
	ms.Push(-2)
	ms.Push(0)
	ms.Push(-3)
	ms.GetMin() // return -3
	ms.Pop()
	ms.Top()           // return 0
	got := ms.GetMin() // return -2
	if got != -2 {
		t.Errorf("want -2 but got %d", got)
	}
}

// Last executed input:
func TestMinStack_Pop2(t *testing.T) {
	ms := &MinStack{}
	ms.Push(0)
	ms.Push(1)
	ms.Push(0)
	ms.GetMin() // return 0
	ms.Pop()
	got := ms.GetMin() // return 0
	if got != 0 {
		t.Errorf("want 0 but got %d", got)
	}
}
