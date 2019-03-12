package main

import "fmt"

type SQueue struct {}

var stack1, stack2 *Stack

func CreateSQueue(capacity int) *SQueue {
	stack1 = createStack(capacity)
	stack2 = createStack(capacity)
	return &SQueue{}
}


func (sq *SQueue) Enqueue(data int) {
	stack1.Push(data)
	stack1.PrintStack()
}

func (sq *SQueue) Dequeue() {
	if stack2.Empty() {
		for !stack1.Empty() {
			stack2.Push(stack1.Top())
			stack1.Pop()
		}
	}
	fmt.Println("Dequeued element: ", stack2.Top())
	stack2.Pop()
}

