package main

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
	size int
}

func CreateLinkedList() *LinkedList {
	return &LinkedList{
		Head: nil,
		size: 0,
	}
}

func (l *LinkedList) Append(value int) {
	if l.Head == nil {
		node := &Node{
			Data: value,
		}
		l.Head = node
		l.Head.Next = nil
		l.size++
	} else {
		cur := l.Head
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = &Node{
			Data: value,
			Next: nil,
		}
		l.size++
	}
}

func (l *LinkedList) InsertAt(index, value int) {
	if index < l.size {
		if index == 0 {
			node := &Node{
				Data: value,
			}
			node.Next = l.Head
			l.Head = node
			l.size++
		} else {
			curIndex := 0
			cur := l.Head
			for curIndex < index - 1 {
				cur = cur.Next
				curIndex++
			}
			node := &Node{
				Data: value,
			}
			node.Next = cur.Next
			cur.Next = node
			l.size++
		}
	} else {
		fmt.Println("\nIndex out of bound")
	}
}

func (l *LinkedList) RemoveAt(index int) {
	if index < l.size {
		if index == 0 {
			cur := l.Head
			l.Head = cur.Next
			cur.Next = nil
			l.size--
		} else {
				curIndex := 0
				node := l.Head
				for curIndex < index - 1{
					node = node.Next
					curIndex++
				}
				cur := node.Next
				node.Next = cur.Next
				cur.Next = nil
				l.size--
		}
	} else {
		fmt.Println("\nIndex out of bound")
	}
}

func (l *LinkedList) PrintLinkedList() {
	if l.size == 0{
		fmt.Println("LinkedList is empty")
	} else {
		node := l.Head
		for len:= 0; len < l.size; len++ {
			fmt.Print(node.Data)
			node = node.Next
		}
	}
}