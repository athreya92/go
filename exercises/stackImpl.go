package main

import "fmt"

type Stack struct {
	n []int
}

func (s *Stack) Empty() {
	fmt.Println(len(s.n) == 0)
}

func (s *Stack) Push(v int) {
	(s.n) = append(s.n, v)
	fmt.Println(s.n)
}

func (s *Stack) Pop()  {
	(s.n) = (s.n)[:len(s.n)-1]
	fmt.Println(s.n)
}
