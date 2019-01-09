package main

func main() {

	Swap(10, 20)
	Swap("shreyas", "athreya")
	Reverse([]interface{}{1, 3, 5, 7, 10})
	Reverse([]interface{}{"I","LOVE", "INDIA", 5, 7, 10})
	Palindrome([]interface{}{"m", "a", "l", "a", "y", "a", "l", "a", "m"})
	CopySlice([]interface{}{3, 4})
	CopySlice([]interface{}{"a", "b"})
	CopyMap(map[interface{}]interface{}{1: "xyz", 2: "php"})

	var s Stack
	s.Empty()
	s.Push(1)
	s.Push(2)
	s.Empty()
	s.Pop()
	s.Pop()
	s.Empty()

	var q Queue
	q.Empty()
	q.Enqueue(3)
	q.Enqueue(4)
	q.Empty()
	q.Dequeue()
	q.Dequeue()
	q.Empty()
}

