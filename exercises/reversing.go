package main

import (
	"fmt"
	"reflect"
)

func Reverse(a []interface{}) {
	for i, j := 0, len(a) - 1; i < j; i, j = i + 1, j - 1 {
		a[i], a[j] = a[j], a[i]
	}
	fmt.Println(a)
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
