package main

import "fmt"

func main() {
	ReverseWords("I am an asshole")
	FizzBuzz(3, 50)
	fibonacci(10)
	fmt.Println(AllUniqueSet("ashole"))
	s := createStack(3)
	s.Push(5)
	s.Push(10)
	s.Push(12)
	s.Push(13)
	s.Top()
	s.Pop()
	s.Top()
	s.FirstElement()

	q := createQueue(3)
	q.Enqueue(10)
	q.Enqueue(25)
	q.Enqueue(100)
	q.Enqueue(235)
	q.First()
	q.Dequeue()
	moveZeroesToEnd([]int{0,0,1,0,3,5,0,8})

	fmt.Println(sortRGB2([]string{"G", "B", "R", "R", "B", "R", "G"}))
	fmt.Println(MultiplyArraySubtractIndex([]int{1, 4, 5, 7, 5}))
	fmt.Println(firstNonDuplicate("qeywerryq"))

}
