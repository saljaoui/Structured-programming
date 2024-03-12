package main

import "fmt"

// Node represents a node in the linked list
type Node struct {
	data interface{}
	next *Node
}

// Stack represents a stack implemented using a linked list
type Stack struct {
	top *Node
	size int
}

// Push adds a new element to the stack
func (s *Stack) Push(value interface{}) {
newNode = &Node{data: value, next: s.top}
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

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}
