package main

import "fmt"

func moveZeroesToEnd(arr []int) {
	i := 0
	for j := 0; j < len(arr); j++ {
		if arr[j] != 0 {
			arr[i] = arr[j]
			i++
			fmt.Println(cap(arr))
		}
	}

	for ; i < len(arr); i++ {
		arr[i] = 0
	}

	fmt.Println(arr)
}

// [1, 2, 0, 3, 0, 5]
// [0, 0, 1, 2, 3, 5]

func moveZeroesToBeginning(arr []int) {
	i := len(arr) - 1
	for j := len(arr) - 1; j >= 0; j-- {
		if arr[j] != 0 {
			arr[i] = arr[j]
			i--
		}
	}
	for ; i >= 0; i-- {
		arr[i] = 0
	}

	fmt.Println(arr)
}

