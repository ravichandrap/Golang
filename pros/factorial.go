package main

import "fmt"

func factorial(fact int) int {
	if fact <= 0 {
		return 1
	}

	return fact * factorial(fact-1)
}

func main() {
	result := factorial(4)
	fmt.Println(result)
}
