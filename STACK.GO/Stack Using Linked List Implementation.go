package main

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

type Stack struct {
	top *Node
	size int
}

func (s *Stack) Push(value interface{}) {
newNode = &Node{data: value, next: s.top}
s.top = newNode
s.size++
}

func (s *Stack) Pop() interface{} {

}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}