package main

import "fmt"

func fibonacci(r int) {
	a := 0
	b := 1
	fmt.Println(a)
	fmt.Println(b)
	for r - 2 > 0 {
		c := a + b
		fmt.Println(c)
		a = b
		b = c
		r--
	}
}
