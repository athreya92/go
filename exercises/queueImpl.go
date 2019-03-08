package main

import "fmt"

type Queue struct {
	n []int
	cap int
}

func createQueue(capacity int) *Queue {
	return &Queue{
		cap: capacity,
		n: []int{},
	}
}

func (q *Queue) Empty() {
	fmt.Println(len(q.n) == 0)
}

func (q *Queue) First() {
	fmt.Println(q.n[0])
}

func (q *Queue) Enqueue(v int) {
	if len(q.n) == q.cap {
		fmt.Println("Queue is full")
	} else {
		q.n = append(q.n, v)
		fmt.Println(q.n)
	}
}

func (q *Queue) Dequeue() {
	q.n = (q.n)[1:len(q.n)]
	fmt.Println(q.n)
}
