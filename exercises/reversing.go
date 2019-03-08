package main

import (
	"fmt"
	"reflect"
)

func ReverseBits(a []byte, s, e int) {
	for e > s {
	a[e], a[s] = a[s], a[e]
	s++
	e--
	}
}

func Palindrome(a []interface{}) {
	check := a
	for i, j := 0, len(a) - 1; i < j; i, j = i + 1, j - 1 {
		a[i], a[j] = a[j], a[i]
	}
	if reflect.DeepEqual(a, check) {
		fmt.Println("Palindrome found")
	} else {
		fmt.Println("Palindrome not found")
	}

}

func ReverseWords(s string) {
	r := []byte(s)
	ReverseBits(r, 0, len(s)-1)

	p := 0
	for q := p; q < len(r) - 1; q++ {
		if r[q] == ' ' {
			ReverseBits(r, p, q -1)
			p = q + 1
		}
	}
	fmt.Println(string(r))

}
