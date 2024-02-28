package main

import "fmt"

type Stack struct {
	items []interface{}
}
func (s *Stack) push(item interface{}) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (interface{}, error) {

}

func main() {






}