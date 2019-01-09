package main

import "fmt"

type Queue struct {
	n []int
}

func (q Queue) Empty() {
	fmt.Println(len(q.n) == 0)
}

func (q *Queue) Enqueue(v int) {
	q.n = append(q.n, v)
	fmt.Println(q.n)
}

func (q *Queue) Dequeue() {
	q.n = (q.n)[1:len(q.n)]
	fmt.Println(q.n)
}
