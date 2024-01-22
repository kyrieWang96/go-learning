package main

import "fmt"

// SequentialStack 顺序栈： 容量固定，顺序栈则需要预留足够的内存空间，以容纳可能的最大栈大小。扩容的话需要重新申请一个大容量并且把原有数据复制到新的栈中
type SequentialStack struct {
	data []any `json:"data"` // 栈存储数据的结构
	top  int   `json:"top"`  // 栈顶 默认为0
	size int   `json:"size"` // 栈容量
}

// NewSequentialStack creates a new SequentialStack
func NewSequentialStack(size int) *SequentialStack {
	return &SequentialStack{
		data: make([]any, size),
		top:  0,
		size: size,
	}
}

// LengthStack returns the number of stack elements
func (s *SequentialStack) LengthStack() int {
	return s.top
}

// IsFull returns true if the stack element is full
func (s *SequentialStack) IsFull() bool {
	return s.top == s.size
}

// IsEmpty returns true if the stack element is empty
func (s *SequentialStack) IsEmpty() bool {
	return s.top == 0
}

// Push Inserts the stack element
func (s *SequentialStack) Push(item any) bool {
	if s.IsFull() {
		fmt.Println("栈已满无法插入对象")
		return false
	}

	s.data[s.top] = item
	s.top++
	return true
}

//Pop get element from stack
func (s *SequentialStack) Pop() any {
	if s.IsEmpty() {
		fmt.Println("栈已空，无数据")
		return nil
	}
	s.top--
	return s.data[s.top]
}

// Peek get element from stack but not pop element
func (s *SequentialStack) Peek() any {
	if s.IsEmpty() {
		fmt.Println("栈已空，无数据")
		return nil
	}
	return s.data[s.top-1]
}
