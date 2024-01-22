package main

import (
	"fmt"
	"testing"
)

func Test_SequentialStackPopData(t *testing.T) {
	stack := NewSequentialStack(10)
	for i := 0; i < 10; i++ {
		if stack.IsFull() {
			t.Errorf("stack is full")
			return
		}
		stack.Push(i)
	}

	fmt.Println(stack.data)

	for !stack.IsEmpty() {
		fmt.Println(stack.Pop())
	}
}
