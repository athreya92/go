package main

import "fmt"

func main() {
	// This will not print since the main function will finish before the goroutine will execute
	go func() {
		fmt.Println("Hello")
	}()
}

func chanProblem() {
	//Functionality of the channel is to block the receiver till the sender passes value to the channel.
	//This will create a deadlock since the channel is not assigned with receiver before sender.
	ch := make(chan int)
	ch <- 42
	fmt.Println(<-ch)
}

func chanSolution1() {
	// Solution 1: Over come the above problem by creating buffered channel.
	ch := make(chan int, 1)
	ch <- 42
	fmt.Println(<-ch)
}

func chanSolution2() {
	// Solution 2: Create a goroutine which will be blocked till it receives the value from the channel
	ch := make(chan int)
	go func(chan int) {
		fmt.Println(<- ch)
	}(ch)
	ch <- 42
}


