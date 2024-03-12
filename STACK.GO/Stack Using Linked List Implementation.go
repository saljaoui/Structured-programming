package main

import (
	"fmt"
)

// Node represents a node in the linked list
type Node struct {
	data interface{}
	next *Node
}

// Stack represents a stack implemented using a linked list
type Stack struct {
	top  *Node
	size int
}

// Push adds a new element to the stack
func (s *Stack) Push(value interface{}) {
	newNode := &Node{data: value, next: s.top}
	s.top = newNode
	s.size++
}

// Pop removes and returns the top element from the stack
func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	data := s.top.data
	s.top = s.top.next
	s.size--
	return data
}

// Peek returns the top element without removing it from the stack
func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.top.data
}

// IsEmpty returns true if the stack is empty
func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

// Size returns the number of elements in the stack
func (s *Stack) Size() int {
	return s.size
}

func main() {
	stack := &Stack{}

	stack.Push(1)
	stack.Push(42)
	stack.Push(55)
	stack.Push(9)
	stack.Push(0)
	stack.Push(23)

	fmt.Println("Stack size:", stack.Size())

	fmt.Println("Top element:", stack.Peek())

	fmt.Println("Pop:", stack.Pop())
	fmt.Println("Pop:", stack.Pop())
	fmt.Println("Pop:", stack.Pop())
	fmt.Println("Pop:", stack.Pop())
	fmt.Println("Pop:", stack.Pop())
	fmt.Println("Pop:", stack.Pop())

	fmt.Println("Is stack empty?", stack.IsEmpty())
}
