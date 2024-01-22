package main

import "testing"

func Test_PushLinkStack(t *testing.T) {
	linkStack := NewLinkListStak()
	for i := 0; i < 10; i++ {
		linkStack.Push(i)
	}

	linkStack.Print()
}
