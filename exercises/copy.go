package main

import "fmt"

func CopySlice(a []interface{}) {
	var b = []interface{}{1, 2}
	check := b
	fmt.Println("value a = ", a)
	fmt.Println("current check = ", check)
	b = a
	fmt.Println("copied b = ", b)
}

func CopyMap(a map[interface{}]interface{}) {
	var b  = make(map[interface{}]interface{})
	fmt.Println("intial b = ", b)
	for key, value := range a {
		b[key] = value
	}
	fmt.Println("copied b = ",b)
}
