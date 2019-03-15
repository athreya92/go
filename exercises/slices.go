package main

import (
	"bytes"
	"fmt"
)

func SliceWindow(arr []byte){
	// reference https://appliedgo.net/slices/
	// slices are windows to under laying arrays. It is just a struct with a pointer to starting of the array, length and capacity
	// type slice struct{
	//	pointer *ptrToArr
	//  len      int
	//  cap      int
	fmt.Println(string(arr))
	s := arr[2:len(arr)-1]
	fmt.Println(string(s))
	xyz := append(arr,'w', 'o', 'r', 'l', 'd')
	fmt.Println(string(xyz))
	fmt.Println(string(arr))

	sp := bytes.Split(arr, []byte(" "))
	fmt.Println(string(sp[0]))

}
