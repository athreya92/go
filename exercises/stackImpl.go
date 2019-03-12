package main

import "fmt"

type Stack struct {
	n []int
	cap int
}

func createStack(capacity int) *Stack {
	return &Stack{
		cap: capacity,
		n: []int{},
	}
}

func (s *Stack) Empty() bool{
	return len(s.n) == 0
}

func (s *Stack) Top() int {
	return s.n[len(s.n) - 1]
}

func (s *Stack) FirstElement() {
	fmt.Println(s.n[0])
}

func (s *Stack) Push(v int) {
	if len(s.n) == s.cap{
		fmt.Println("Stack is full")
	} else {
		s.n = append(s.n, v)
	}
}

func (s *Stack) Pop() {
	s.n = (s.n)[:len(s.n)-1]
}

func (s *Stack) PrintStack() {
	for _, v := range s.n {
		fmt.Println(v)
	}
}