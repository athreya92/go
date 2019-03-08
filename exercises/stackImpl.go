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

func (s *Stack) Empty() {
	fmt.Println(len(s.n) == 0)
}

func (s *Stack) Top() {
	fmt.Println(s.n[len(s.n) - 1])
}

func (s *Stack) FirstElement() {
	fmt.Println(s.n[0])
}

func (s *Stack) Push(v int) {
	if len(s.n) == s.cap{
		fmt.Println("Stack is full")
	} else {
		s.n = append(s.n, v)
		fmt.Println(s.n)
	}
}

func (s *Stack) Pop()  {
	s.n = (s.n)[0:len(s.n)-1]
	fmt.Println(s.n)
}