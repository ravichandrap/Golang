package main

import "fmt"

func factorial(fact int) int {
	if fact <= 1 {
		return 1
	}

	return fact * factorial(fact-1)
}

func main() {
	result := factorial(4)
	fmt.Println(result)
	//example: 4 = 4 * 3 * 2 * 1 -> 24
}
