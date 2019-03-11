package main

import "fmt"

type LQueue struct {
	head *Node
	ptr *Node
	size int
}

func CreateLQueue() *LQueue {
	return &LQueue{
		head: nil,
		ptr: nil,
		size: 0,
	}
}

func (lq *LQueue) Enqueue(data int) {
	if lq.head == nil {
		node := &Node{
			Data:data,
		}
		lq.head = node
		lq.head.Next = nil
		lq.ptr = lq.head
		lq.size++
	} else {
		node := &Node{
			Data:data,
		}
		cur := lq.ptr
		lq.ptr = node
		cur.Next = lq.ptr
		lq.size++
	}
}

func (lq *LQueue) Dequeue() {
	if lq.head == nil {
		fmt.Println("Queue is empty")
	} else {
		cur := lq.head
		lq.head = cur.Next
		cur = nil
		lq.size--
	}
}

func (lq *LQueue) PrintQueue() {
	cur := lq.head

	for len := 0; len < lq.size; len ++ {
		fmt.Println(cur.Data)
		cur = cur.Next
	}
}
