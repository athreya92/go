package main

import (
	"fmt"
)

type LStack struct {
	top *Node
	size int
}

func CreateLStack() *LStack {
	return &LStack{
		top: nil,
		size: 0,
	}
}

func (ls *LStack) Push(data int) {
	if ls.top == nil {
		node := &Node{
			Data: data,
		}
		ls.top = node
		ls.top.Next = nil
		ls.size++
	} else {
		node := &Node{
			Data: data,
		}
		cur := ls.top
		ls.top = node
		ls.top.Next = cur
		ls.size++
	}
}

func (ls *LStack) Pop() {
	if ls.top == nil {
		fmt.Println("Stack is empty")
	} else {
		ls.top = ls.top.Next
		ls.size--
	}
}

func (ls *LStack) PrintStack() {
	cur := ls.top

	for len := 0; len < ls.size; len ++ {
		fmt.Print("\t",  cur.Data)
		cur = cur.Next
	}
}