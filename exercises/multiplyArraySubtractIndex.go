package main

func MultiplyArraySubtractIndex(input []int) []int {
	total := 1

	for i := 0; i < len(input); i++ {
		if input[i] > 0 {
			total *= input[i]
		}
	}

	for i := len(input) - 1; i >= 0; i-- {
		if input[i] > 0 {
			input[i] = total / input[i]
		} else {
			input[i] = total
		}
	}

	return input
}
